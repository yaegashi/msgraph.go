package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/yaegashi/msgraph.go/auth"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

const (
	defaultTenantID = "common"
	defaultClientID = "45c7f99c-0a94-42ff-a6d8-a8d657229e8c"
	defaultScope    = "openid profile user.readwrite files.readwrite.all"
)

func dump(o interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(o)
}

func main() {
	var tenantID, clientID string
	flag.StringVar(&tenantID, "tenant_id", defaultTenantID, "Tenant ID")
	flag.StringVar(&clientID, "client_id", defaultClientID, "Client ID")
	flag.Parse()

	da, err := auth.NewDeviceAuth(tenantID, clientID, defaultScope)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(da.Message())

	dt, err := da.Poll()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	cli := dt.Client(ctx)
	serv := msgraph.NewService(cli)

	{
		s := serv.Me()
		log.Printf("GET %s", s.URL())
		r, err := s.DoGet()
		if err == nil {
			dump(r)
		} else {
			log.Println(err)
		}
	}

	{
		s := serv.Me().Drive().Root().Children()
		q := "?" + url.Values{"$filter": {"file ne null"}, "$select": {"name,file,size,webUrl"}}.Encode()
		log.Printf("GET %s%s", s.URL(), q)
		r, err := s.DoGetWithPath(q)
		if err == nil {
			dump(r)
		} else {
			log.Println(err)
		}
	}
}
