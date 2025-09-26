package api

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ArtyomKr/tg-spotify-bot/internal/storage"
)

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Healthy\n")
}

func (s *Server) HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	userID := r.URL.Query().Get("state")

	err := s.storage.Set(userID, storage.UserData{Code: code})
	if err != nil {
		log.Printf("Failed to save user data for %s: %v", userID, err)
		http.Error(w, "An internal error occurred. Please try again later.", http.StatusInternalServerError)
		return
	}

	redirectUrl := os.Getenv("TG_BOT_LINK")

	http.Redirect(w, r, redirectUrl, http.StatusFound)
}
