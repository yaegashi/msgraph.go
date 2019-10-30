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
	defaultScope    = "Sites.Read.All Files.Read.All"
)

// Config is serializable configuration
type Config struct {
	TenantID     string `json:"tenant_id,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	TokenStore   string `json:"token_store,omitempty"`
}

// App is application
type App struct {
	Config
	URL         string
	Token       *auth.Token
	GraphClient *msgraph.GraphServiceRequestBuilder
}

// Authenticate performs OAuth2 authentication
func (app *App) Authenticate(ctx context.Context) error {
	var err error
	m := auth.NewTokenManager()
	if app.ClientSecret == "" {
		if app.TokenStore != "" {
			// ignore errors
			m.Load(app.TokenStore)
		}
		app.Token, err = m.DeviceAuthorizationGrant(app.TenantID, app.ClientID, defaultScope, nil)
		if err != nil {
			return err
		}
		if app.TokenStore != "" {
			err = m.Save(app.TokenStore)
			if err != nil {
				return err
			}
		}
	} else {
		app.Token, err = m.ClientCredentialsGrant(app.TenantID, app.ClientID, app.ClientSecret, auth.DefaultMSGraphScope)
		if err != nil {
			return err
		}
	}
	app.GraphClient = msgraph.NewClient(app.Token.Client(ctx))
	return nil
}

// Main is main routine
func (app *App) Main(ctx context.Context) error {
	if app.URL == "" {
		return fmt.Errorf("URL not specified")
	}
	err := app.Authenticate(ctx)
	if err != nil {
		return err
	}
	itemRB, err := app.GraphClient.GetDriveItemByURL(ctx, app.URL)
	if err != nil {
		return err
	}
	itemRB.Workbook().Worksheets()
	app.GraphClient.Workbooks()
	item, err := itemRB.Request().Get(ctx)
	if err != nil {
		return err
	}
	log.Printf("Download to %#v (%d bytes)", *item.Name, *item.Size)
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
		return nil
	}
	return fmt.Errorf("unknown download URL")
}

func main() {
	config := ""
	app := &App{}
	flag.StringVar(&config, "config", "", "Config file path")
	flag.StringVar(&app.TenantID, "tenant-id", "", "Tenant ID (default:"+defaultTenantID+")")
	flag.StringVar(&app.ClientID, "client-id", "", "Client ID (default: "+defaultClientID+")")
	flag.StringVar(&app.ClientSecret, "client-secret", "", "Client secret (for client credentials grant)")
	flag.StringVar(&app.TokenStore, "token-store", "", "OAuth2 token store path")
	flag.StringVar(&app.URL, "url", "", "URL")
	flag.Parse()

	cfg := &Config{
		TenantID: defaultTenantID,
		ClientID: defaultClientID,
	}
	if config != "" {
		b, err := ioutil.ReadFile(config)
		if err != nil {
			log.Fatal(err)
		}
		err = jsonx.Unmarshal(b, cfg)
		if err != nil {
			log.Fatal(err)
		}
	}
	if app.TenantID == "" {
		app.TenantID = cfg.TenantID
	}
	if app.ClientID == "" {
		app.ClientID = cfg.ClientID
	}
	if app.ClientSecret == "" {
		app.ClientSecret = cfg.ClientSecret
	}
	if app.TokenStore == "" {
		app.TokenStore = cfg.TokenStore
	}

	ctx := context.Background()
	err := app.Main(ctx)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
