package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	msgraph "github.com/yaegashi/msgraph.go/beta"
	"github.com/yaegashi/msgraph.go/jsonx"
	"github.com/yaegashi/msgraph.go/msauth"
	P "github.com/yaegashi/msgraph.go/ptr"
	"golang.org/x/oauth2"
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
		Password:                      P.String(uuid.New().String()),
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
	enc := jsonx.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(o)
}

func main() {
	var tenantID, clientID, clientSecret string
	flag.StringVar(&tenantID, "tenant-id", defaultTenantID, "Tenant ID")
	flag.StringVar(&clientID, "client-id", defaultClientID, "Client ID")
	flag.StringVar(&clientSecret, "client-secret", defaultClientSecret, "Client Secret")
	flag.Parse()

	ctx := context.Background()
	m := msauth.NewManager()
	scopes := []string{msauth.DefaultMSGraphScope}
	ts, err := m.ClientCredentialsGrant(ctx, tenantID, clientID, clientSecret, scopes)
	if err != nil {
		log.Fatal(err)
	}

	httpClient := oauth2.NewClient(ctx, ts)
	graphClient := msgraph.NewClient(httpClient)

	{
		// Search and delete default user
		r := graphClient.Users().Request()
		r.Filter(fmt.Sprintf("userPrincipalName eq '%s'", *defaultUser.UserPrincipalName))
		users, _ := r.Get(ctx)
		for _, user := range users {
			log.Printf("Delete user %s", *user.ID)
			graphClient.Users().ID(*user.ID).Request().Delete(ctx)
		}
	}

	{
		// Search and delete default group
		r := graphClient.Groups().Request()
		r.Filter(fmt.Sprintf("mailNickname eq '%s'", *defaultGroup.MailNickname))
		groups, _ := r.Get(ctx)
		for _, group := range groups {
			log.Printf("Delete group %s", *group.ID)
			graphClient.Groups().ID(*group.ID).Request().Delete(ctx)
		}
	}

	var u *msgraph.User
	{
		log.Printf("Create user %s", *defaultUser.UserPrincipalName)
		u, err = graphClient.Users().Request().Add(ctx, defaultUser)
		if err != nil {
			log.Fatal(err)
		}
		dump(u)
	}

	var g *msgraph.Group
	{
		log.Printf("Create group %s", *defaultGroup.DisplayName)
		g, err = graphClient.Groups().Request().Add(ctx, defaultGroup)
		if err != nil {
			log.Fatal(err)
		}
		dump(g)
	}

	{
		log.Printf("Add member %s to group %s", *u.ID, *g.ID)
		reqObj := map[string]interface{}{
			"@odata.id": graphClient.DirectoryObjects().ID(*u.ID).Request().URL(),
		}
		r := graphClient.Groups().ID(*g.ID).Members().Request()
		err := r.JSONRequest(ctx, "POST", "/$ref", reqObj, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	{
		log.Printf("Show user %s", *u.ID)
		r := graphClient.Users().ID(*u.ID).Request()
		r.Expand("memberOf")
		users, err := r.Get(ctx)
		if err != nil {
			log.Fatal(err)
		}
		dump(users)
	}

	{
		log.Printf("Show group %s", *g.ID)
		r := graphClient.Groups().ID(*g.ID).Request()
		r.Expand("members")
		groups, err := r.Get(ctx)
		if err != nil {
			log.Fatal(err)
		}
		dump(groups)
	}
}
