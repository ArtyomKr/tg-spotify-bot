package spotify

import (
	"net/url"
	"os"
)

func GetLoginLink(state string) string {
	const scopes = "user-read-playback-state user-modify-playback-state"
	const adress = "https://accounts.spotify.com/authorize"
	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	redirectURI := os.Getenv("SPOTIFY_REDIRECT_URI")

	query := url.Values{
		"response_type": {"code"},
		"client_id":     {clientId},
		"scope":         {scopes},
		"redirect_uri":  {redirectURI},
		"state":         {state},
	}

	link := adress + "?" + query.Encode()
	return link
}
