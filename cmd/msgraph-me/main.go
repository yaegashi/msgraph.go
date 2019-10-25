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

	"github.com/yaegashi/msgraph.go/auth"
	"github.com/yaegashi/msgraph.go/jsonx"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

const (
	defaultTenantID = "common"
	defaultClientID = "45c7f99c-0a94-42ff-a6d8-a8d657229e8c"
	defaultScope    = "openid profile offline_access User.Read Files.Read"
	tokenStorePath  = "token_store.json"
)

func dump(o interface{}) {
	enc := jsonx.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(o)
}

func main() {
	var tenantID, clientID string
	flag.StringVar(&tenantID, "tenant_id", defaultTenantID, "Tenant ID")
	flag.StringVar(&clientID, "client_id", defaultClientID, "Client ID")
	flag.Parse()

	m := auth.NewTokenManager()
	m.Load(tokenStorePath)
	dt, err := m.DeviceAuthorizationGrant(tenantID, clientID, defaultScope, nil)
	if err != nil {
		log.Fatal(err)
	}
	m.Save(tokenStorePath)

	ctx := context.Background()
	cli := dt.Client(ctx)
	serv := msgraph.NewClient(cli)

	{
		log.Printf("Get current logged in user information")
		r := serv.Me().Request()
		log.Printf("GET %s", r.URL())
		x, err := r.Get()
		if err == nil {
			dump(x)
		} else {
			log.Println(err)
		}
	}

	var items []msgraph.DriveItem
	{
		log.Printf("Get files in the root folder of user's drive")
		r := serv.Me().Drive().Root().Children().Request()
		r.Filter("file ne null")
		r.Select("id,name,file,size,webUrl")
		log.Printf("GET %s", r.URL())
		items, err = r.Get()
		if err == nil {
			dump(items)
		} else {
			log.Println(err)
		}
	}

	for _, item := range items {
		err := func() error {
			r := serv.Me().Drive().Items().ID(*item.ID).Request()
			log.Printf("GET %s", r.URL())
			x, err := r.Get()
			if err != nil {
				return err
			}
			log.Printf("Download to %#v (%d bytes)", *item.Name, *item.Size)
			if url, ok := x.GetAdditionalData("@microsoft.graph.downloadUrl"); ok {
				res, err := http.Get(url.(string))
				if err != nil {
					return err
				}
				defer res.Body.Close()
				if res.StatusCode != http.StatusOK {
					b, _ := ioutil.ReadAll(res.Body)
					return fmt.Errorf("%s: %s", res.Status, string(b))
				}
				f, err := os.Create(*x.Name)
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
