package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/oauth2"
)

const (
	endpointURLFormat         = "https://login.microsoftonline.com/%s/oauth2/v2.0/%s"
	deviceCodeGrantType       = "urn:ietf:params:oauth:grant-type:device_code"
	refresTokenGrantType      = "refresh_token"
	authorizationPendingError = "authorization_pending"
)

// DeviceCode is returned on device auth inintiation
type DeviceCode struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURL string `json:"verification_url"`
	ExpiresIn       int    `json:"expires_in"`
	Interval        int    `json:"interval"`
	Message         string `json:"message"`
}

// DeviceToken is returned on device auth success
type DeviceToken struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiresOn    int    `json:"expires_on"`
	AccessToken  string `json:"access_token"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}

// DeviceError is returned on errors
type DeviceError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// DeviceTokenManager is a persistent store for DeviceToken
type DeviceTokenManager struct {
	Store map[string]*DeviceToken
	Dirty bool
}

// NewDeviceTokenManager returns a new DeviceTokenManager instance
func NewDeviceTokenManager() *DeviceTokenManager {
	return &DeviceTokenManager{Store: map[string]*DeviceToken{}}
}

// Load loads DeviceToken store
func (m *DeviceTokenManager) Load(path string) error {
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

// Save saves DeviceToken store
func (m *DeviceTokenManager) Save(path string) error {
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

// Authenticate acquires DeviceToken
func (m *DeviceTokenManager) Authenticate(tenantID, clientID, scope string, callback func(*DeviceCode) error) (*DeviceToken, error) {
	key := fmt.Sprintf("%s:%s", tenantID, clientID)
	devicecodeURL := fmt.Sprintf(endpointURLFormat, tenantID, "devicecode")
	tokenURL := fmt.Sprintf(endpointURLFormat, tenantID, "token")
	dt, ok := m.Store[key]
	if !ok {
		dt = &DeviceToken{}
	}
	now := int(time.Now().Unix())
	if now < dt.ExpiresOn {
		return dt, nil
	}
	for dt.RefreshToken != "" {
		values := url.Values{
			"client_id":     {clientID},
			"grant_type":    {refresTokenGrantType},
			"scope":         {scope},
			"refresh_token": {dt.RefreshToken},
		}
		res, err := http.PostForm(tokenURL, values)
		if err != nil {
			break
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			break
		}
		dt = &DeviceToken{}
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(dt)
		if err != nil {
			break
		}
		if dt.ExpiresOn == 0 {
			dt.ExpiresOn = int(time.Now().Unix()) + dt.ExpiresIn
		}
		m.Store[key] = dt
		m.Dirty = true
		return dt, nil
	}
	res, err := http.PostForm(devicecodeURL, url.Values{"client_id": {clientID}, "scope": {scope}})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("%s: %s", res.Status, string(b))
	}
	dc := &DeviceCode{}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&dc)
	if err != nil {
		return nil, err
	}
	err = callback(dc)
	if err != nil {
		return nil, err
	}
	values := url.Values{
		"client_id":   {clientID},
		"grant_type":  {deviceCodeGrantType},
		"device_code": {dc.DeviceCode},
	}
	interval := dc.Interval
	if interval == 0 {
		interval = 5
	}
	for {
		time.Sleep(time.Second * time.Duration(interval))
		res, err := http.PostForm(tokenURL, values)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		if res.StatusCode == http.StatusOK {
			dt := &DeviceToken{}
			dec := json.NewDecoder(res.Body)
			err = dec.Decode(dt)
			if err != nil {
				return nil, err
			}
			if dt.ExpiresOn == 0 {
				dt.ExpiresOn = int(time.Now().Unix()) + dt.ExpiresIn
			}
			m.Store[key] = dt
			m.Dirty = true
			return dt, nil
		}
		de := &DeviceError{}
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(de)
		if err != nil {
			return nil, err
		}
		if de.Error != authorizationPendingError {
			return nil, fmt.Errorf("%s: %s", de.Error, de.ErrorDescription)
		}
	}
}

// Token implements oauth2.TokenSource interface
func (dt *DeviceToken) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		TokenType:    dt.TokenType,
		AccessToken:  dt.AccessToken,
		RefreshToken: dt.RefreshToken,
	}, nil
}

// Client returns *http.Client
func (dt *DeviceToken) Client(ctx context.Context) *http.Client {
	return oauth2.NewClient(ctx, dt)
}
