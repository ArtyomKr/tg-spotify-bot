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

	currentTrack, err := b.spotifyAPI.GetCurrentTrack(token)
	if err != nil {
		log.Printf("Req error: %v", err)
		return tgbotapi.NewMessage(msg.Chat.ID, "Couldn't get playback status")
	}

	var messageBuilder strings.Builder

	messageBuilder.WriteString("<b>🎵 Now Playing:</b>\n")
	messageBuilder.WriteString(fmt.Sprintf("<a href=\"%s\"><b>%s</b></a>\n", currentTrack.Item.ExternalUrls.Spotify, currentTrack.Item.Name))
	messageBuilder.WriteString(fmt.Sprintf("by <i>%s</i>\n", currentTrack.Item.Artists[0].Name))
	messageBuilder.WriteString(fmt.Sprintf("from <a href=\"%s\"><i>%s</i></a>\n", currentTrack.Item.Album.ExternalUrls.Spotify, currentTrack.Item.Album.Name))

	messageBuilder.WriteString(formatPlaybackProgress(currentTrack.ProgressMs, currentTrack.Item.DurationMs))

	response := tgbotapi.NewMessage(msg.Chat.ID, messageBuilder.String())
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

	var messageBuilder strings.Builder

	if len(queue.Queue) > 0 {
		messageBuilder.WriteString("<b>📊 Queue:</b>\n")
		for i, track := range queue.Queue {
			messageBuilder.WriteString(fmt.Sprintf("%d. <b>%s</b>\n", i+1, track.Name))
			messageBuilder.WriteString(fmt.Sprintf("   by <i>%s</i>\n", track.Artists[0].Name))
			messageBuilder.WriteString("\n")
		}
	}

	response := tgbotapi.NewMessage(msg.Chat.ID, messageBuilder.String())
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
