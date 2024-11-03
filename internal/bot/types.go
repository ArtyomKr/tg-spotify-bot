package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/internal/auth"
	"telegram-bot/internal/spotify"
	"telegram-bot/internal/storage"
)

type Bot struct {
	api         *tgbotapi.BotAPI
	spotifyAPI  *spotify.Client
	storage     storage.UserStorage
	spotifyAuth *auth.Manager
}

type CommandHandler interface {
	Handle(msg *tgbotapi.Message) tgbotapi.MessageConfig
}
