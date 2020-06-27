package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/yaegashi/msgraph.go/jsonx"
	"github.com/yaegashi/msgraph.go/msauth"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
	"github.com/yaegashi/wtz.go"
	"golang.org/x/oauth2"
)

const (
	defaultTenantID       = "common"
	defaultClientID       = "45c7f99c-0a94-42ff-a6d8-a8d657229e8c"
	defaultTokenCachePath = "token_cache.json"
)

var defaultScopes = []string{"offline_access", "User.Read", "Calendars.Read", "Files.Read"}

func dump(o interface{}) {
	enc := jsonx.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(o)
}

func main() {
	var tenantID, clientID, tokenCachePath string
	flag.StringVar(&tenantID, "tenant-id", defaultTenantID, "Tenant ID")
	flag.StringVar(&clientID, "client-id", defaultClientID, "Client ID")
	flag.StringVar(&tokenCachePath, "token-cache-path", defaultTokenCachePath, "Token cache path")
	flag.Parse()

	ctx := context.Background()
	m := msauth.NewManager()
	m.LoadFile(tokenCachePath)
	ts, err := m.DeviceAuthorizationGrant(ctx, tenantID, clientID, defaultScopes, nil)
	if err != nil {
		log.Fatal(err)
	}
	m.SaveFile(tokenCachePath)

	httpClient := oauth2.NewClient(ctx, ts)
	graphClient := msgraph.NewClient(httpClient)

	{
		log.Printf("Get current logged in user information")
		req := graphClient.Me().Request()
		log.Printf("GET %s", req.URL())
		user, err := req.Get(ctx)
		if err == nil {
			dump(user)
		} else {
			log.Println(err)
		}
	}

	{
		log.Println("Get current logged in user's first 10 events")
		req := graphClient.Me().Events().Request()
		req.Top(10)
		tzname, err := wtz.LocationToName(time.Local)
		if err != nil {
			log.Println(err)
			tzname = "Tokyo Standard Time"
		}
		req.Header().Add("Prefer", fmt.Sprintf(`outlook.timezone="%s"`, tzname))
		log.Printf("GET %s", req.URL())
		events, err := req.GetN(ctx, 1)
		if err == nil {
			dump(events)
		} else {
			log.Println(err)
		}
	}

	var items []msgraph.DriveItem
	{
		log.Printf("Get files in the root folder of user's drive")
		req := graphClient.Me().Drive().Root().Children().Request()
		// This filter is not supported by OneDrive for Business or SharePoint Online
		// https://docs.microsoft.com/en-us/onedrive/developer/rest-api/concepts/filtering-results?#filterable-properties
		// req.Filter("file ne null")
		log.Printf("GET %s", req.URL())
		items, err = req.Get(ctx)
		if err != nil {
			log.Println(err)
		}
	}

	for _, item := range items {
		timestamp := item.LastModifiedDateTime.Format(time.RFC3339)
		itemType := "FILE"
		if item.File == nil {
			itemType = "DIR "
		}
		log.Printf("  %s %s %10d %s", itemType, timestamp, *item.Size, *item.Name)
	}

	fmt.Print("Press ENTER to download files or Ctrl-C to abort: ")
	fmt.Scanln()

	for _, item := range items {
		if item.File == nil {
			continue
		}
		err := func() error {
			log.Printf("Download %q (%d bytes)", *item.Name, *item.Size)
			if url, ok := item.GetAdditionalData("@microsoft.graph.downloadUrl"); ok {
				res, err := http.Get(url.(string))
				if err != nil {
					return err
				}
				defer res.Body.Close()
				if res.StatusCode != http.StatusOK {
					b, _ := ioutil.ReadAll(res.Body)
					return fmt.Errorf("%s: %s", res.Status, string(b))
				}
				f, err := os.Create(*item.Name)
				if err != nil {
					return err
				}
				defer f.Close()
				_, err = io.Copy(f, res.Body)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("unknown download URL")
			}
			return nil
		}()
		if err != nil {
			log.Println(err)
		}
	}
}
