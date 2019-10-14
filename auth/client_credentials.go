package auth

import "net/url"

const (
	clientCredentialsGrantType = "client_credentials"
)

// AuthenticateClientCredentials authenticates via OAuth2 client credentials flow
func (m *TokenManager) AuthenticateClientCredentials(tenantID, clientID, clientSecret, scope string) (*Token, error) {
	t, err := m.refreshToken(tenantID, clientID, clientSecret, scope)
	if err == nil && t != nil {
		return t, nil
	}
	values := url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"scope":         {scope},
		"grant_type":    {clientCredentialsGrantType},
	}
	return m.requestToken(tenantID, clientID, values)
}
