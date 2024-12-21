package bot

import (
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
		return b.handleQueue(msg)
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

	playbackStatus, err := b.spotifyAPI.GetCurrentTrack(token)
	if err != nil {
		log.Printf("Req error: %v", err)
		return tgbotapi.NewMessage(msg.Chat.ID, "Couldn't get playback status")
	}

	message := formatCurrentlyPlaying(playbackStatus)

	response := tgbotapi.NewMessage(msg.Chat.ID, message)
	response.ParseMode = "HTML"

	return response
}

func (b *Bot) handleQueue(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	id := strconv.FormatInt(msg.From.ID, 10)
	token, err := b.spotifyAuth.GetToken(id)
	if err != nil {
		return tgbotapi.NewMessage(msg.Chat.ID, "Authorization failed")
	}

	queue, err := b.spotifyAPI.GetQueue(token)
	if err != nil {
		log.Printf("Req error: %v", err)
		return tgbotapi.NewMessage(msg.Chat.ID, "Couldn't get playback status")
	}

	message := formatQueueItems(&queue.Queue)

	response := tgbotapi.NewMessage(msg.Chat.ID, message)
	response.ParseMode = "HTML"

	return response
}

func (b *Bot) addToQueue(msg *tgbotapi.Message, uris []string) tgbotapi.MessageConfig {
	id := strconv.FormatInt(msg.From.ID, 10)
	token, err := b.spotifyAuth.GetToken(id)
	if err != nil {
		return tgbotapi.NewMessage(msg.Chat.ID, "Authorization failed")
	}

	for _, uri := range uris {
		err = b.spotifyAPI.AddTrackToQueue(token, uri)
		if err != nil {
			log.Printf("Req error: %v", err)
			return tgbotapi.NewMessage(msg.Chat.ID, "Couldn't add track to queue")
		}
	}

	successText := "Added track to queue"
	if len(uris) > 1 {
		successText = "Added tracks to queue"
	}

	return tgbotapi.NewMessage(msg.Chat.ID, successText)
}

func (b *Bot) handleUnknownCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	uris := getSpotifyURI(msg.Text)
	if len(uris) != 0 {
		return b.addToQueue(msg, uris)
	}
	return tgbotapi.NewMessage(msg.Chat.ID, "Unknown command")
}
