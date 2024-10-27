package server

import (
	"io"
	"net/http"
	"os"
	"telegram-bot/internal/storage"
)

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Healthy\n")
}

func (s *Server) HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	userID := r.URL.Query().Get("state")

	s.storage.Set(userID, storage.UserData{Code: code})

	redirectUrl := os.Getenv("TG_BOT_LINK")

	http.Redirect(w, r, redirectUrl, http.StatusFound)
}
