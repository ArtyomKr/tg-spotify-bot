package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"telegram-bot/internal/bot"
	"telegram-bot/internal/server"
	"telegram-bot/internal/storage"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Couldn't load env variables")
	}

	port := os.Getenv("PORT")
	token := os.Getenv("TG_BOT_TOKEN")

	userStorage := storage.NewStorage()

	srv := server.New(port, userStorage)
	srv.Listen()

	tgbot, err := bot.NewBot(token, userStorage)
	if err != nil {
		log.Fatal(err)
	}

	tgbot.Start()
}
