package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"telegram-bot/internal/message"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Couldn't load env variables")
	}

	token := os.Getenv("TG_BOT_TOKEN")
	log.Printf("Starting the bot with token %v", token)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf(update.Message.CommandArguments())
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := message.Process(update.Message)
			//msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
