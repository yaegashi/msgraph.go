package adal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest/adal"
	"golang.org/x/oauth2"
)

// Client is Microsoft Graph Client
type Client struct {
	Authority   string
	TenantID    string
	ClientID    string
	Resource    string
	Store       string
	oauthConfig *adal.OAuthConfig
	spToken     *adal.ServicePrincipalToken
}

// LoadToken loads tokens from the persistent store.  If it's unavailable, run device code authentication.
func (cli *Client) LoadToken() error {
	callback := func(token adal.Token) error { return cli.SaveToken(token) }
	var err error
	cli.oauthConfig, err = adal.NewOAuthConfig(cli.Authority, cli.TenantID)
	if err != nil {
		return err
	}
	token, err := adal.LoadToken(cli.Store)
	if err == nil {
		spToken, err := adal.NewServicePrincipalTokenFromManualToken(*cli.oauthConfig, cli.ClientID, cli.Resource, *token, callback)
		if err == nil {
			spToken.SetAutoRefresh(true)
			if spToken.EnsureFresh() == nil {
				cli.spToken = spToken
				return nil
			}
		}
	}
	oauthClient := &http.Client{}
	deviceCode, err := adal.InitiateDeviceAuth(oauthClient, *cli.oauthConfig, cli.ClientID, cli.Resource)
	if err != nil {
		return err
	}
	fmt.Println(*deviceCode.Message)
	token, err = adal.WaitForUserCompletion(oauthClient, deviceCode)
	if err != nil {
		return err
	}
	spToken, err := adal.NewServicePrincipalTokenFromManualToken(*cli.oauthConfig, cli.ClientID, cli.Resource, *token, callback)
	if err != nil {
		return err
	}
	spToken.SetAutoRefresh(true)
	cli.SaveToken(spToken.Token())
	cli.spToken = spToken
	return nil
}

// SaveToken saves OAuth2 token in the persistent store.
func (cli *Client) SaveToken(token adal.Token) error {
	return adal.SaveToken(cli.Store, 0600, token)
}

// NewHTTPClient returns http.Client with OAuth2 authorization.
func (cli *Client) NewHTTPClient(ctx context.Context) *http.Client {
	token := cli.spToken.OAuthToken()
	return oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))
}
