package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (b *Bot) addTrackToQueue(msg *tgbotapi.Message, trackURIs []string) tgbotapi.MessageConfig {
	id := strconv.FormatInt(msg.From.ID, 10)
	token, err := b.spotifyAuth.GetToken(id)
	if err != nil {
		return tgbotapi.NewMessage(msg.Chat.ID, "Authorization failed")
	}

	for _, uri := range trackURIs {
		err = b.spotifyAPI.AddTrackToQueue(token, uri)
		if err != nil {
			log.Printf("Req error: %v", err)
			return tgbotapi.NewMessage(msg.Chat.ID, "Couldn't add track to queue")
		}
	}

	successText := "Added track to queue"
	if len(trackURIs) > 1 {
		successText = "Added tracks to queue"
	}

	return tgbotapi.NewMessage(msg.Chat.ID, successText)
}

func (b *Bot) addAlbumToQueue(msg *tgbotapi.Message, albumID string) tgbotapi.MessageConfig {
	id := strconv.FormatInt(msg.From.ID, 10)
	token, err := b.spotifyAuth.GetToken(id)
	if err != nil {
		return tgbotapi.NewMessage(msg.Chat.ID, "Authorization failed")
	}

	tracks, err := b.spotifyAPI.GetAlbumTracks(token, albumID)
	if err != nil {
		return tgbotapi.NewMessage(msg.Chat.ID, "Could not get album")
	}

	trackURIs := make([]string, len(tracks.Items))
	for i, track := range tracks.Items {
		trackURIs[i] = track.URI
	}

	if len(trackURIs) != 0 {
		return b.addTrackToQueue(msg, trackURIs)
	}

	return tgbotapi.NewMessage(msg.Chat.ID, "Couldn't add this album to queue")
}
