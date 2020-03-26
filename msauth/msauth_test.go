package msauth_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/yaegashi/msgraph.go/msauth"
	"golang.org/x/oauth2"
)

const (
	tenantID       = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
	clientID       = "YYYYYYYY-YYYY-YYYY-YYYY-YYYYYYYYYYYY"
	clientSecret   = "ZZZZZZZZZZZZZZZZZZZZZZZZ"
	username       = "user.name@your-domain.com"
	password       = "secure-password"
	tokenStorePath = "token_store.json"
)

var (
	daScopes = []string{"openid", "profile", "offline_access", "User.Read", "Files.Read"}
	ccScopes = []string{msauth.DefaultMSGraphScope}
)

func ExampleManager_DeviceAuthorizationGrant() {
	ctx := context.Background()
	m := msauth.NewManager()
	m.LoadFile(tokenStorePath)
	ts, err := m.DeviceAuthorizationGrant(ctx, tenantID, clientID, daScopes, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = m.SaveFile(tokenStorePath)
	if err != nil {
		log.Fatal(err)
	}
	httpClient := oauth2.NewClient(ctx, ts)
	res, err := httpClient.Get("https://graph.microsoft.com/v1.0/me")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(string(b)))
}

func ExampleManager_ClientCredentialsGrant() {
	ctx := context.Background()
	m := msauth.NewManager()
	ts, err := m.ClientCredentialsGrant(ctx, tenantID, clientID, clientSecret, ccScopes)
	if err != nil {
		log.Fatal(err)
	}
	httpClient := oauth2.NewClient(ctx, ts)
	res, err := httpClient.Get("https://graph.microsoft.com/v1.0/me")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(string(b)))
}

func ExampleManager_ResourceOwnerPasswordGrant() {
	ctx := context.Background()
	m := msauth.NewManager()
	ts, err := m.ResourceOwnerPasswordGrant(ctx, tenantID, clientID, clientSecret, username, password, ccScopes)
	if err != nil {
		log.Fatal(err)
	}
	httpClient := oauth2.NewClient(ctx, ts)
	res, err := httpClient.Get("https://graph.microsoft.com/v1.0/me")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(string(b)))
}
