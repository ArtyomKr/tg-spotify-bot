package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (b *Bot) handleMessage(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	switch msg.Text {
	case "/login":
		return b.handleLogin(msg)
	case "/token":
		return b.handleToken(msg)
	case "/current":
		return b.handleCurrent(msg)

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
		return tgbotapi.NewMessage(msg.Chat.ID, "Couldn't get playback status")
	}

	return tgbotapi.NewMessage(msg.Chat.ID, currentTrack.Item.Name)
}

func (b *Bot) handleUnknownCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(msg.Chat.ID, "Unknown command")
}
