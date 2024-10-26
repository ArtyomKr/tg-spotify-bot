package server

import (
	"log"
	"net/http"
)

type UserData struct {
	code  string
	token string
}

type Server struct {
	port string
	data map[string]UserData
}

func New(port string) *Server {
	return &Server{port: port}
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
