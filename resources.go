package goteams

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// GenerateToken get the token to communicate with the bot
// It returns the token type and the access token
func GenerateToken(clientID, clientSecret string) (string, error) {
	urlLogin := "https://login.microsoftonline.com/botframework.com/oauth2/v2.0/token"

	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	data.Add("client_id", clientID)
	data.Add("client_secret", clientSecret)
	data.Add("scope", "https://api.botframework.com/.default")

	req, _ := http.NewRequest("POST", urlLogin, strings.NewReader(data.Encode()))
	req.Header.Add("Host", "login.microsoftonline.com")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var oauth2 OAuth2Token
	err = json.NewDecoder(resp.Body).Decode(&oauth2)
	if err != nil {
		return "", err
	}

	token := []string{oauth2.TokenType, oauth2.AccessToken}

	return strings.Join(token, " "), nil
}

// PostMessage sends a personal/channel message to conversation on Teams
func (msg *SendMessage) PostMessage(accessToken string) (int, error) {
	urlApi := fmt.Sprintf("https://smba.trafficmanager.net/amer/v3/conversations/%s/activities/%s", msg.Conversation.ID, msg.ReplyToID)

	data, _ := json.Marshal(msg)
	req, _ := http.NewRequest("POST", urlApi, bytes.NewBuffer(data))
	req.Header.Set("Authorization", fmt.Sprintf("%s", accessToken))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}
