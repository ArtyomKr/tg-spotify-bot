package spotify

import (
	"encoding/json"
	"log"
	"net/http"
)

const playerUrl = "https://api.spotify.com/v1/me/player"

func (c *Client) GetCurrentTrack(token string) (SpotifyPlaybackStatusRes, error) {
	req, err := http.NewRequest("GET", playerUrl+"/currently-playing", nil)
	if err != nil {
		log.Printf("Req error: %v", err)
		return SpotifyPlaybackStatusRes{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Sending error: %v", err)
		return SpotifyPlaybackStatusRes{}, err
	}
	defer res.Body.Close()

	var playbackStatusRes SpotifyPlaybackStatusRes
	if err := json.NewDecoder(res.Body).Decode(&playbackStatusRes); err != nil {
		log.Printf("Response error: %v", err)
		return SpotifyPlaybackStatusRes{}, err
	}

	return playbackStatusRes, nil
}
