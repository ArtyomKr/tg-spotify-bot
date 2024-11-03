package spotify

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) GetLoginLink(state string) string {
	const scopes = "user-read-playback-state user-modify-playback-state"
	const adress = "https://accounts.spotify.com/authorize"

	query := url.Values{
		"response_type": {"code"},
		"client_id":     {c.clientID},
		"scope":         {scopes},
		"redirect_uri":  {c.redirectURI},
		"state":         {state},
	}

	link := adress + "?" + query.Encode()
	return link
}

func (c *Client) GetToken(code string) (TokenResBody, error) {
	authCode := base64.StdEncoding.EncodeToString([]byte(c.clientID + ":" + c.clientSecret))

	formData := url.Values{}
	formData.Set("grant_type", "authorization_code")
	formData.Set("code", code)
	formData.Set("redirect_uri", c.redirectURI)

	req, err := http.NewRequest("POST", c.apiURL+"/token", strings.NewReader(formData.Encode()))
	if err != nil {
		log.Printf("Req error: %v", err)
		return TokenResBody{}, err
	}

	req.Header.Set("Authorization", "Basic "+authCode)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Sending error: %v", err)
		return TokenResBody{}, err
	}
	defer res.Body.Close()

	var tokenRes TokenResBody
	if err := json.NewDecoder(res.Body).Decode(&tokenRes); err != nil {
		log.Printf("Response error: %v", err)
		return TokenResBody{}, err
	}

	return tokenRes, nil
}

func (c *Client) RefreshToken(refreshToken string) (TokenResBody, error) {
	authCode := base64.StdEncoding.EncodeToString([]byte(c.clientID + ":" + c.clientSecret))

	formData := url.Values{}
	formData.Set("refresh_token", refreshToken)
	formData.Set("grant_type", "refresh_token")

	req, err := http.NewRequest("POST", c.apiURL+"/token", strings.NewReader(formData.Encode()))
	if err != nil {
		log.Printf("Req error: %v", err)
		return TokenResBody{}, err
	}

	req.Header.Set("Authorization", "Basic "+authCode)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Sending error: %v", err)
		return TokenResBody{}, err
	}
	defer res.Body.Close()

	var tokenRes TokenResBody
	if err := json.NewDecoder(res.Body).Decode(&tokenRes); err != nil {
		log.Printf("Response error: %v", err)
		return TokenResBody{}, err
	}

	if tokenRes.RefreshToken == "" {
		tokenRes.RefreshToken = refreshToken
	}

	return tokenRes, nil
}
