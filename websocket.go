package main

import (
	"log"
	"net/http"

	"github.com/olahol/melody"
)

func handleWebSockets(path string) {
	m := melody.New()

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	m.HandleMessage(func(s *melody.Session, b []byte) {
		log.Println(string(b))
		m.Broadcast(b)
	})
}
