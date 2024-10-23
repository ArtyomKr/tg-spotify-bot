package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panic("Couldn't load env variables")
	}

	fmt.Println(GetLoginLink())

}

func GetLoginLink() string {
	const scopes = "user-read-playback-state user-modify-playback-state"
	const adress = "https://accounts.spotify.com/authorize"
	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	redirectURI := os.Getenv("SPOTIFY_REDIRECT_URI")

	query := fmt.Sprintf("?response_type=%s&client_id=%s&scope=%s&redirect_uri=%s",
		url.QueryEscape("code"), url.QueryEscape(clientId), url.QueryEscape(scopes), url.QueryEscape(redirectURI))

	link := adress + query
	return link
}
