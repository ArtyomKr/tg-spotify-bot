package main

import (
	"log"
	"os"

	"github.com/ArtyomKr/tg-spotify-bot/internal/api"
	"github.com/ArtyomKr/tg-spotify-bot/internal/auth"
	"github.com/ArtyomKr/tg-spotify-bot/internal/bot"
	"github.com/ArtyomKr/tg-spotify-bot/internal/spotify"
	"github.com/ArtyomKr/tg-spotify-bot/internal/storage"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Couldn't find env variables in .env file")
	}

	port := os.Getenv("PORT")
	token := os.Getenv("TG_BOT_TOKEN")

	userStorage, err := storage.NewSqlLiteStorage("./users.db")
	if err != nil {
		log.Panic("Couldn't create storage file")
	}
	defer userStorage.Close()

	spotifyClient := spotify.NewClient()
	spotifyAuth := auth.NewManager(userStorage, spotifyClient)
	srv := api.New(port, userStorage)

	srv.Listen()
	tgbot, err := bot.NewBot(token, userStorage, spotifyClient, spotifyAuth)
	if err != nil {
		log.Fatal(err)
	}

	tgbot.Start()
}
