package bot

import (
	"github.com/ArtyomKr/tg-spotify-bot/internal/auth"
	"github.com/ArtyomKr/tg-spotify-bot/internal/spotify"
	"github.com/ArtyomKr/tg-spotify-bot/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
