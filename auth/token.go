package auth

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	// DefaultMSGraphScope is the default scope for MS Graph API
	DefaultMSGraphScope = "https://graph.microsoft.com/.default"
)

// Token contains OAuth2 tokens returned by the token endpoint
type Token struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiresOn    int    `json:"expires_on"`
	AccessToken  string `json:"access_token"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}

// Token implements oauth2.TokenSource interface
func (t *Token) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		TokenType:    t.TokenType,
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
	}, nil
}

// Client returns *http.Client
func (t *Token) Client(ctx context.Context) *http.Client {
	return oauth2.NewClient(ctx, t)
}

// TokenError is returned on failed authentication
type TokenError struct {
	ErrorX           string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// Error implements error interface
func (t *TokenError) Error() string {
	return fmt.Sprintf("%s: %s", t.ErrorX, t.ErrorDescription)
}
