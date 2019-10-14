package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	endpointURLFormat     = "https://login.microsoftonline.com/%s/oauth2/v2.0/%s"
	refreshTokenGrantType = "refresh_token"
)

func generateKey(tenantID, clientID string) string {
	return fmt.Sprintf("%s:%s", tenantID, clientID)
}

func deviceCodeURL(tenantID string) string {
	return fmt.Sprintf(endpointURLFormat, tenantID, "devicecode")
}

func tokenURL(tenantID string) string {
	return fmt.Sprintf(endpointURLFormat, tenantID, "token")
}

// TokenManager is a persistent store for Token
type TokenManager struct {
	Store map[string]*Token
	Dirty bool
}

// NewTokenManager returns a new TokenManager instance
func NewTokenManager() *TokenManager {
	return &TokenManager{Store: map[string]*Token{}}
}

// Load loads Token store
func (m *TokenManager) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	err = dec.Decode(&m.Store)
	if err != nil {
		return err
	}
	m.Dirty = false
	return nil
}

// Save saves Token store
func (m *TokenManager) Save(path string) error {
	if !m.Dirty {
		return nil
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(file)
	err = enc.Encode(m.Store)
	if err != nil {
		file.Close()
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	m.Dirty = false
	return nil
}

func (m *TokenManager) requestToken(tenantID, clientID string, values url.Values) (*Token, error) {
	res, err := http.PostForm(tokenURL(tenantID), values)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		t := &Token{}
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(t)
		if err != nil {
			return nil, err
		}
		if t.ExpiresOn == 0 {
			t.ExpiresOn = int(time.Now().Unix()) + t.ExpiresIn
		}
		m.Store[generateKey(tenantID, clientID)] = t
		m.Dirty = true
		return t, nil
	}
	t := &TokenError{}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(t)
	if err != nil {
		return nil, err
	}
	return nil, t
}

func (m *TokenManager) refreshToken(tenantID, clientID, clientSecret, scope string) (*Token, error) {
	t, ok := m.Store[generateKey(tenantID, clientID)]
	if !ok {
		t = &Token{}
	}
	if int(time.Now().Unix()) < t.ExpiresOn {
		return t, nil
	}
	if t.RefreshToken == "" {
		return nil, fmt.Errorf("No refresh token")
	}
	values := url.Values{
		"client_id":     {clientID},
		"grant_type":    {refreshTokenGrantType},
		"scope":         {scope},
		"refresh_token": {t.RefreshToken},
	}
	if clientSecret != "" {
		values.Add("client_secret", clientSecret)
	}
	return m.requestToken(tenantID, clientID, values)
}
