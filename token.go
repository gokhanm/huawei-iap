package huaweistore

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
)

const (
	// ProductTokenURL is the endpoint for production environment.
	ProductTokenURL string = "https://oauth-login.cloud.huawei.com/oauth2/v2/token"
)

// GetHeaderToken get access token
func (c *Client) GetHeaderToken(reqBody IAPRequest) (string, error) {
	appAt, err := c.getToken(reqBody)
	if err != nil {
		return "", err
	}

	oriString := fmt.Sprintf("APPAT:%s", appAt)
	auth := base64.StdEncoding.EncodeToString([]byte(oriString))
	authHeader := fmt.Sprintf("Basic %s", auth)

	return authHeader, nil
}

func (c *Client) getToken(reqBody IAPRequest) (string, error) {
	urlValue := url.Values{
		"grant_type":    {"client_credentials"},
		"client_secret": {reqBody.ClientSecret},
		"client_id":     {reqBody.ClientID},
	}

	resp, err := c.httpCli.PostForm(c.TokenURL, urlValue)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var token TokenResponse
	if err := json.Unmarshal(bodyBytes, &token); err != nil {
		return "", err
	}

	if token.AccessToken == "" {
		return "", errors.New("Get token fail, " + string(bodyBytes))
	}

	return token.AccessToken, nil

}
