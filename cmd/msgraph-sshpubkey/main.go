package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/yaegashi/msgraph.go/jsonx"
	"github.com/yaegashi/msgraph.go/msauth"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
	"golang.org/x/oauth2"
)

const (
	defaultExtensionName = "dev.l0w.ssh_public_keys"
	defaultTenantID      = "common"
	defaultClientID      = "45c7f99c-0a94-42ff-a6d8-a8d657229e8c"
)

var defaultScopes = []string{"openid", "User.ReadWrite"}

// Config is serializable configuration
type Config struct {
	TenantID      string            `json:"tenant_id,omitempty"`
	ClientID      string            `json:"client_id,omitempty"`
	ClientSecret  string            `json:"client_secret,omitempty"`
	ExtensionName string            `json:"extension_name,omitempty"`
	TokenCache    string            `json:"token_cache,omitempty"`
	LoginMap      map[string]string `json:"login_map"`
}

// App is application
type App struct {
	Config
	Op          string
	In          string
	Out         string
	Login       string
	TokenSource oauth2.TokenSource
	GraphClient *msgraph.GraphServiceRequestBuilder
}

// User returns target graph user endpoint
func (app *App) User() *msgraph.UserRequestBuilder {
	if app.Login == "" {
		return app.GraphClient.Me()
	}
	userID := app.Login
	if app.LoginMap != nil {
		if id, ok := app.LoginMap[app.Login]; ok {
			userID = id
		}
	}
	return app.GraphClient.Users().ID(userID)
}

// Authenticate performs OAuth2 authentication
func (app *App) Authenticate(ctx context.Context) error {
	var err error
	m := msauth.NewManager()
	if app.ClientSecret == "" {
		if app.TokenCache != "" {
			// ignore errors
			m.LoadFile(app.TokenCache)
		}
		app.TokenSource, err = m.DeviceAuthorizationGrant(ctx, app.TenantID, app.ClientID, defaultScopes, nil)
		if err != nil {
			return err
		}
		if app.TokenCache != "" {
			err = m.SaveFile(app.TokenCache)
			if err != nil {
				return err
			}
		}
	} else {
		scopes := []string{msauth.DefaultMSGraphScope}
		app.TokenSource, err = m.ClientCredentialsGrant(ctx, app.TenantID, app.ClientID, app.ClientSecret, scopes)
		if err != nil {
			return err
		}
	}
	app.GraphClient = msgraph.NewClient(oauth2.NewClient(ctx, app.TokenSource))
	return nil
}

// Set performs set operation on user's extension
func (app *App) Set(ctx context.Context) error {
	err := app.Authenticate(ctx)
	if err != nil {
		return err
	}
	var in []byte
	if app.In == "-" {
		in, err = ioutil.ReadAll(os.Stdin)
	} else {
		in, err = ioutil.ReadFile(app.In)
	}
	if err != nil {
		return err
	}
	newExt := &msgraph.Extension{}
	newExt.SetAdditionalData("extensionName", app.ExtensionName)
	newExt.SetAdditionalData("value", string(in))
	_, err = app.User().Extensions().Request().Add(ctx, newExt)
	if err != nil {
		if errRes, ok := err.(*msgraph.ErrorResponse); ok {
			if errRes.StatusCode() == http.StatusConflict {
				err = app.User().Extensions().ID(app.ExtensionName).Request().Update(ctx, newExt)
			}
		}
	}
	return err
}

// Get performs get operation on user's extensions
func (app *App) Get(ctx context.Context) error {
	err := app.Authenticate(ctx)
	if err != nil {
		return err
	}
	r := app.User().Request()
	r.Select("id")
	r.Expand("extensions")
	user, err := r.Get(ctx)
	if err != nil {
		return err
	}
	for _, x := range user.Extensions {
		if *x.ID != app.ExtensionName {
			continue
		}
		value, _ := x.GetAdditionalData("value")
		if s, ok := value.(string); ok {
			if app.Out == "-" {
				fmt.Print(s)
				return nil
			}
			return ioutil.WriteFile(app.Out, []byte(s), 0644)
		}
		return fmt.Errorf("No value in extension %s", app.ExtensionName)
	}
	return nil
}

// Delete performs delete operation on user's extension
func (app *App) Delete(ctx context.Context) error {
	err := app.Authenticate(ctx)
	if err != nil {
		return err
	}
	return app.User().Extensions().ID(app.ExtensionName).Request().Delete(ctx)
}

func main() {
	config := ""
	app := &App{}
	flag.StringVar(&config, "config", "", "Config file path")
	flag.StringVar(&app.Op, "op", "GET", "Operation (GET/SET/DELETE)")
	flag.StringVar(&app.TenantID, "tenant-id", "", "Tenant ID (default:"+defaultTenantID+")")
	flag.StringVar(&app.ClientID, "client-id", "", "Client ID (default: "+defaultClientID+")")
	flag.StringVar(&app.ClientSecret, "client-secret", "", "Client secret (for client credentials grant)")
	flag.StringVar(&app.TokenCache, "token-cache", "", "OAuth2 token cache path")
	flag.StringVar(&app.ExtensionName, "extension-name", "", "Extension name (default: "+defaultExtensionName+")")
	flag.StringVar(&app.Login, "login", "", "Login name (default: authenticated user)")
	flag.StringVar(&app.In, "in", "", "Input file (\"-\" for stdin)")
	flag.StringVar(&app.Out, "out", "-", "Output file (\"-\" for stdout)")
	flag.Parse()

	cfg := &Config{
		TenantID:      defaultTenantID,
		ClientID:      defaultClientID,
		ExtensionName: defaultExtensionName,
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
	if app.TokenCache == "" {
		app.TokenCache = cfg.TokenCache
	}
	if app.ExtensionName == "" {
		app.ExtensionName = cfg.ExtensionName
	}
	app.LoginMap = cfg.LoginMap

	var err error
	ctx := context.Background()
	switch strings.ToLower(app.Op) {
	case "set":
		err = app.Set(ctx)
	case "get":
		err = app.Get(ctx)
	case "delete":
		err = app.Delete(ctx)
	default:
		flag.CommandLine.SetOutput(os.Stderr)
		flag.Usage()
		os.Exit(1)
	}
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
