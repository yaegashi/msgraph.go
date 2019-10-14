package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	deviceCodeGrantType       = "urn:ietf:params:oauth:grant-type:device_code"
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

// AuthenticateDeviceAuth authenticates via OAuth2 device auth flow
func (m *TokenManager) AuthenticateDeviceAuth(tenantID, clientID, scope string, callback func(*DeviceCode) error) (*Token, error) {
	t, err := m.refreshToken(tenantID, clientID, "", scope)
	if err == nil && t != nil {
		return t, nil
	}
	res, err := http.PostForm(deviceCodeURL(tenantID), url.Values{"client_id": {clientID}, "scope": {scope}})
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
		t, err := m.requestToken(tenantID, clientID, values)
		if err == nil {
			return t, nil
		}
		tokenError, ok := err.(*TokenError)
		if !ok || tokenError.ErrorX != authorizationPendingError {
			return nil, err
		}
	}
}
