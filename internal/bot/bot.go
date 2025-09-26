package bot

import (
	"log"

	"github.com/ArtyomKr/tg-spotify-bot/internal/auth"
	"github.com/ArtyomKr/tg-spotify-bot/internal/spotify"
	"github.com/ArtyomKr/tg-spotify-bot/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewBot(token string, storage storage.UserStorage, spotifyAPI *spotify.Client, spotifyAuth *auth.Manager) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api,
		spotifyAPI,
		storage,
		spotifyAuth,
	}, nil
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.api.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg := b.handleMessage(update.Message)
			//msg.ReplyToMessageID = update.Message.MessageID
			b.api.Send(msg)
		}
	}
}
