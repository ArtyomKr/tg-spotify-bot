package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"telegram-bot/internal/spotify"
)

func Process(msg *tgbotapi.Message) tgbotapi.MessageConfig {
	switch msg.Text {
	case "/login":
		id := strconv.FormatInt(msg.From.ID, 10)
		return tgbotapi.NewMessage(msg.Chat.ID, spotify.GetLoginLink(id))
	}
	return tgbotapi.NewMessage(msg.Chat.ID, "Unknown command")
}
