package goteams

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// GenerateToken get the token to communicate with the bot
// It returns the token type and the access token
func GenerateToken(client_id, client_secret string) (string, string, error) {
	urlLogin := "https://login.microsoftonline.com/botframework.com/oauth2/v2.0/token"

	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	data.Add("client_id", client_id)
	data.Add("client_secret", client_secret)
	data.Add("scope", "https://api.botframework.com/.default")

	req, _ := http.NewRequest("POST", urlLogin, strings.NewReader(data.Encode()))
	req.Header.Add("Host", "login.microsoftonline.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	var token OAuth2Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", "", err
	}

	return token.TokenType, token.AccessToken, nil
}