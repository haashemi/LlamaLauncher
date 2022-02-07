package egl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const AccountURL = "https://account-public-service-prod.ol.epicgames.com/account/api/oauth"

type GrantedToken struct {
	AccessToken      string        `json:"access_token"`
	ExpiresIn        int           `json:"expires_in"`
	ExpiresAt        time.Time     `json:"expires_at"`
	TokenType        string        `json:"token_type"`
	RefreshToken     string        `json:"refresh_token"`
	RefreshExpires   int           `json:"refresh_expires"`
	RefreshExpiresAt time.Time     `json:"refresh_expires_at"`
	AccountID        string        `json:"account_id"`
	ClientID         string        `json:"client_id"`
	InternalClient   bool          `json:"internal_client"`
	ClientService    string        `json:"client_service"`
	Scope            []interface{} `json:"scope"`
	DisplayName      string        `json:"displayName"`
	App              string        `json:"app"`
	InAppID          string        `json:"in_app_id"`
	DeviceID         string        `json:"device_id"`
}

type ExchangeData struct {
	ExpiresInSeconds int    `json:"expiresInSeconds"`
	Code             string `json:"code"`
	CreatingClientId string `json:"creatingClientId"`
}

func grantToken(payloads map[string]string, client Client, userAgent string) (*GrantedToken, error) {
	requestPayloads := url.Values{}
	for key, value := range payloads {
		requestPayloads.Set(key, value)
	}

	req, err := http.NewRequest(http.MethodPost, AccountURL+"/token", bytes.NewReader([]byte(requestPayloads.Encode())))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "basic "+client.Encode())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, FindError(resp.Body)
	}
	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := &GrantedToken{}
	return data, json.Unmarshal(raw, data)
}

func OAuthFromAuthorizationCode(code string, client Client, userAgent string) (*GrantedToken, error) {
	return grantToken(map[string]string{
		"grant_type": "authorization_code",
		"code":       code,
	}, client, userAgent)
}

func OAuthFromRefreshToken(refreshToken string, client Client, userAgent string) (*GrantedToken, error) {
	return grantToken(map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}, client, userAgent)
}

func GetExchange(accessToken string, client Client, userAgent string) (*ExchangeData, error) {
	req, err := http.NewRequest(http.MethodGet, AccountURL+"/exchange", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "bearer "+accessToken)
	req.Header.Set("User-Agent", userAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, FindError(resp.Body)
	}
	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := &ExchangeData{}
	return data, json.Unmarshal(raw, data)
}
