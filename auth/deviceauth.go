package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"
)

const (
	DeviceCodeGrantType = "urn:ietf:params:oauth:grant-type:device_code"
)

type DeviceAuth struct {
	Code     *DeviceCode
	TenantID string
	ClientID string
}

type DeviceCode struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURL string `json:"verification_url"`
	ExpiresIn       int    `json:"expires_in"`
	Interval        int    `json:"interval"`
	Message         string `json:"message"`
}

type DeviceToken struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}

type DeviceError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func NewDeviceAuth(tenantID, clientID, scope string) (*DeviceAuth, error) {
	devicecodeURL := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/devicecode", tenantID)
	res, err := http.PostForm(devicecodeURL, url.Values{"client_id": {clientID}, "scope": {scope}})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("%s: %s", res.Status, string(b))
	}
	da := &DeviceAuth{ClientID: clientID, TenantID: tenantID}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&da.Code)
	if err != nil {
		return nil, err
	}
	return da, nil
}

func (da *DeviceAuth) Message() string {
	return da.Code.Message
}

func (da *DeviceAuth) Poll() (*DeviceToken, error) {
	tokenURL := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", da.TenantID)
	values := url.Values{
		"client_id":   {da.ClientID},
		"grant_type":  {DeviceCodeGrantType},
		"device_code": {da.Code.DeviceCode},
	}
	interval := da.Code.Interval
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
			return dt, nil
		}
		de := &DeviceError{}
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(de)
		if err != nil {
			return nil, err
		}
		if de.Error != "authorization_pending" {
			return nil, fmt.Errorf("%s: %s", de.Error, de.ErrorDescription)
		}
	}
}

func (dt *DeviceToken) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		TokenType:    dt.TokenType,
		AccessToken:  dt.AccessToken,
		RefreshToken: dt.RefreshToken,
	}, nil
}

func (dt *DeviceToken) Client(ctx context.Context) *http.Client {
	return oauth2.NewClient(ctx, dt)
}
