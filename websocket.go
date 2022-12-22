package main

import (
	"log"
	"net/http"

	"github.com/olahol/melody"
)

var ws *melody.Melody

func handleWebSockets(path string) {
	ws = melody.New()

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		ws.HandleRequest(w, r)
	})

	ws.HandleMessage(func(s *melody.Session, b []byte) {
		log.Println(string(b))
		ws.Broadcast(b)
	})
}
