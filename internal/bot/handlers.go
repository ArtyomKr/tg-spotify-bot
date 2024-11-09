package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func (b *Bot) handleMessage(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	args := strings.Fields(msg.Text)
	if len(args) == 0 {
		return b.handleUnknownCommand(msg)
	}

	command := args[0]

	switch command {
	case "/login":
		return b.handleLogin(msg)
	case "/token":
		return b.handleToken(msg)
	case "/current":
		return b.handleCurrent(msg)
	case "/queue":
		if len(args) < 2 {
			return tgbotapi.NewMessage(msg.Chat.ID, "Please provide a URL. Usage: /queue <url>")
		}
		uri, ok := getSpotifyURI(args[1])
		if !ok {
			return tgbotapi.NewMessage(msg.Chat.ID, "URL is not a valid spotify track url")
		}

		return b.handleQueue(msg, uri)
	default:
		return b.handleUnknownCommand(msg)
	}
}

func (b *Bot) handleLogin(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	id := strconv.FormatInt(msg.From.ID, 10)
	loginLink := b.spotifyAPI.GetLoginLink(id)
	return tgbotapi.NewMessage(msg.Chat.ID, loginLink)
}

func (b *Bot) handleToken(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	id := strconv.FormatInt(msg.From.ID, 10)
	token, err := b.spotifyAuth.GetToken(id)
	//if err != nil {
	//	return tgbotapi.NewMessage(msg.Chat.ID, "You are not authorized with spotify")
	//} // TODO: handle unauthorized requests
	if err != nil {
		return tgbotapi.NewMessage(msg.Chat.ID, "Authorization failed")
	}

	return tgbotapi.NewMessage(msg.Chat.ID, token)
}

func (b *Bot) handleCurrent(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	id := strconv.FormatInt(msg.From.ID, 10)
	token, err := b.spotifyAuth.GetToken(id)
	if err != nil {
		return tgbotapi.NewMessage(msg.Chat.ID, "Authorization failed")
	}

	currentTrack, err := b.spotifyAPI.GetCurrentTrack(token)
	if err != nil {
		log.Printf("Req error: %v", err)
		return tgbotapi.NewMessage(msg.Chat.ID, "Couldn't get playback status")
	}

	response := fmt.Sprintf(
		"⏸️ %s - %s\n⌛️Progress: %s\n%s",
		currentTrack.Item.Name,
		currentTrack.Item.Artists[0].Name,
		formatPlaybackProgress(currentTrack.ProgressMs, currentTrack.Item.DurationMs),
		currentTrack.Item.ExternalUrls.Spotify,
	)
	return tgbotapi.NewMessage(msg.Chat.ID, response)
}

func (b *Bot) handleQueue(msg *tgbotapi.Message, uri string) tgbotapi.MessageConfig {
	id := strconv.FormatInt(msg.From.ID, 10)
	token, err := b.spotifyAuth.GetToken(id)
	if err != nil {
		return tgbotapi.NewMessage(msg.Chat.ID, "Authorization failed")
	}

	err = b.spotifyAPI.AddTrackToQueue(token, uri)
	if err != nil {
		log.Printf("Req error: %v", err)
		return tgbotapi.NewMessage(msg.Chat.ID, "Couldn't add track to queue")
	}

	return tgbotapi.NewMessage(msg.Chat.ID, "Added track to queue")
}

func (b *Bot) handleUnknownCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	uri, ok := getSpotifyURI(msg.Text)
	if ok {
		return b.handleQueue(msg, uri)
	}
	return tgbotapi.NewMessage(msg.Chat.ID, "Unknown command")
}
