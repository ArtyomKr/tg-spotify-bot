package spotify

import (
	"net/http"
	"os"
)

func NewClient() *Client {
	return &Client{
		httpClient:   &http.Client{},
		clientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		clientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		redirectURI:  os.Getenv("SPOTIFY_REDIRECT_URI"),
		apiURL:       os.Getenv("SPOTIFY_API_URL"),
	}
}
