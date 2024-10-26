package server

import (
	"fmt"
	"io"
	"net/http"
)

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Healthy\n")
}

func (s *Server) HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	fmt.Printf("Request data code: %v, state: %v \n", code, state)
}
