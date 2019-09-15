package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"msgraph.go/adal"

	msgraph "msgraph.go/v1.0"
)

const (
	defaultAuthority = "https://login.microsoftonline.com/"
	defaultTenantID  = "common"
	defaultClientID  = "45c7f99c-0a94-42ff-a6d8-a8d657229e8c"
	defaultResource  = "https://graph.microsoft.com"
	defaultStore     = "token_cache.json"
)

func dump(o interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(o)
}

func main() {
	adalClient := &adal.Client{}
	flag.StringVar(&adalClient.Authority, "authority", defaultAuthority, "Authority")
	flag.StringVar(&adalClient.TenantID, "tenant_id", defaultTenantID, "Tenant ID")
	flag.StringVar(&adalClient.ClientID, "client_id", defaultClientID, "Client ID")
	flag.StringVar(&adalClient.Resource, "resource", defaultResource, "Resource")
	flag.StringVar(&adalClient.Store, "store", defaultStore, "Token cache store path")
	flag.Parse()

	err := adalClient.LoadToken()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	cli := adalClient.NewHTTPClient(ctx)
	serv := msgraph.NewService(cli)

	{
		s := serv.Me()
		log.Printf("GET %s", s.URL())
		r, err := s.Get()
		if err == nil {
			dump(r)
		} else {
			log.Println(err)
		}
	}

	{
		s := serv.Me().Drive().Root().Children()
		log.Printf("GET %s", s.URL())
		r, err := s.Get()
		if err == nil {
			dump(r)
		} else {
			log.Println(err)
		}
	}
}
