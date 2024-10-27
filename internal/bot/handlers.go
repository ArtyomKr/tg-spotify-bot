package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (b *Bot) handleMessage(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	switch msg.Text {
	case "/login":
		return b.handleLogin(msg)
	case "/token":
		return b.handleToken(msg)
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
	userData, ok := b.storage.Get(id)
	if !ok {
		return tgbotapi.NewMessage(msg.Chat.ID, "You are not authorized with spotify")
	}

	tokenData, err := b.spotifyAPI.GetToken(userData.Code)
	if err != nil {
		log.Printf("Error: %v", err)
		return tgbotapi.NewMessage(msg.Chat.ID, "Failed to get token")
	}
	return tgbotapi.NewMessage(msg.Chat.ID, tokenData.AccessToken+" "+tokenData.TokenType+""+strconv.Itoa(tokenData.ExpiresIn)+"\n")
}

func (b *Bot) handleUnknownCommand(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(msg.Chat.ID, "Unknown command")
}
