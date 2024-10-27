package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/internal/spotify"
	"telegram-bot/internal/storage"
)

type Bot struct {
	api        *tgbotapi.BotAPI
	spotifyAPI *spotify.Client
	storage    storage.UserStorage
}

type CommandHandler interface {
	Handle(msg *tgbotapi.Message) tgbotapi.MessageConfig
}
