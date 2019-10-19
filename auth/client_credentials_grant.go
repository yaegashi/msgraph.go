package auth

import "net/url"

const (
	clientCredentialsGrantType = "client_credentials"
)

// ClientCredentialsGrant authenticates via OAuth 2.0 client credentials grant
func (m *TokenManager) ClientCredentialsGrant(tenantID, clientID, clientSecret, scope string) (*Token, error) {
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
