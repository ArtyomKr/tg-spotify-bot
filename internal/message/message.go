package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/internal/spotify"
)

func Process(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	switch msg.Text {
	case "/login":
		return tgbotapi.NewMessage(msg.Chat.ID, spotify.GetLoginLink())
	}
	return tgbotapi.NewMessage(msg.Chat.ID, "Unknown command")
}
