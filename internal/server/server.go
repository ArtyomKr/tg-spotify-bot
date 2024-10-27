package server

import (
	"log"
	"net/http"
	"telegram-bot/internal/storage"
)

type Server struct {
	port    string
	storage storage.UserStorage
}

func New(port string, storage storage.UserStorage) *Server {
	return &Server{port, storage}
}

func (s *Server) Listen() {
	http.HandleFunc("/", s.Health)
	http.HandleFunc("/callback", s.HandleCallback)

	go func() {
		err := http.ListenAndServe(":"+s.port, nil)
		if err != nil {
			log.Panic("Couldn't start the server")
		}
	}()

	log.Printf("Server started on port %s", s.port)
}
