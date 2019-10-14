package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yaegashi/msgraph.go/auth"
	msgraph "github.com/yaegashi/msgraph.go/beta"
	P "github.com/yaegashi/msgraph.go/ptr"
)

// Default ID and secret: you should replace them with your own
const (
	defaultTenantID     = "XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
	defaultClientID     = "YYYYYYYY-YYYY-YYYY-YYYY-YYYYYYYYYYYY"
	defaultClientSecret = "ZZZZZZZZZZZZZZZZZZZZZZZZ"
)

// Default user: you should replace it with your own
var defaultUser = &msgraph.User{
	AccountEnabled:    P.Bool(true),
	UserPrincipalName: P.String("hoge-moge@l0w.dev"),
	MailNickname:      P.String("hoge-moge"),
	DisplayName:       P.String("Hoge Moge"),
	GivenName:         P.String("Hoge"),
	Surname:           P.String("Moge"),
	PasswordProfile: &msgraph.PasswordProfile{
		ForceChangePasswordNextSignIn: P.Bool(false),
		Password:                      P.String("XXX!111@YYY#222$ZZZ%333^"),
	},
}

// Default group: you should replace it with your own
var defaultGroup = &msgraph.Group{
	SecurityEnabled: P.Bool(true),
	MailEnabled:     P.Bool(false),
	MailNickname:    P.String("fuga"),
	DisplayName:     P.String("Fuga"),
}

func dump(o interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(o)
}

func main() {
	var tenantID, clientID, clientSecret string
	flag.StringVar(&tenantID, "tenant_id", defaultTenantID, "Tenant ID")
	flag.StringVar(&clientID, "client_id", defaultClientID, "Client ID")
	flag.StringVar(&clientSecret, "client_secret", defaultClientSecret, "Client Secret")
	flag.Parse()

	m := auth.NewTokenManager()
	t, err := m.AuthenticateClientCredentials(tenantID, clientID, clientSecret, auth.DefaultMSGraphScope)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	httpClient := t.Client(ctx)
	graphClient := msgraph.NewClient(httpClient)

	{
		// Search and delete default user
		r := graphClient.Users().Request()
		r.Filter(fmt.Sprintf("userPrincipalName eq '%s'", *defaultUser.UserPrincipalName))
		users, _ := r.Get()
		for _, user := range users {
			log.Printf("Delete user %s", *user.Id)
			graphClient.Users().ID(*user.Id).Request().Delete()
		}
	}

	{
		// Search and delete default group
		r := graphClient.Groups().Request()
		r.Filter(fmt.Sprintf("mailNickname eq '%s'", *defaultGroup.MailNickname))
		groups, _ := r.Get()
		for _, group := range groups {
			log.Printf("Delete group %s", *group.Id)
			graphClient.Groups().ID(*group.Id).Request().Delete()
		}
	}

	var u *msgraph.User
	{
		log.Printf("Create user %s", *defaultUser.UserPrincipalName)
		u, err = graphClient.Users().Request().Add(defaultUser)
		if err != nil {
			log.Fatal(err)
		}
		dump(u)
	}

	var g *msgraph.Group
	{
		log.Printf("Create group %s", *defaultGroup.DisplayName)
		g, err = graphClient.Groups().Request().Add(defaultGroup)
		if err != nil {
			log.Fatal(err)
		}
		dump(g)
	}

	{
		log.Printf("Add member %s to group %s", *u.Id, *g.Id)
		reqObj := map[string]interface{}{
			"@odata.id": graphClient.DirectoryObjects().ID(*u.Id).Request().URL(),
		}
		r := graphClient.Groups().ID(*g.Id).Members().Request()
		err := r.JSONRequestWithPath("POST", "/$ref", reqObj, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	{
		log.Printf("Show user %s", *u.Id)
		r := graphClient.Users().ID(*u.Id).Request()
		r.Expand("memberOf")
		users, err := r.Get()
		if err != nil {
			log.Fatal(err)
		}
		dump(users)
	}

	{
		log.Printf("Show group %s", *g.Id)
		r := graphClient.Groups().ID(*g.Id).Request()
		r.Expand("members")
		groups, err := r.Get()
		if err != nil {
			log.Fatal(err)
		}
		dump(groups)
	}
}
