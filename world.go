package main

import (
	"net/http"
	"time"

	"github.com/olahol/melody"
)

var w bouncyballs

func handleWebSockets(path string) {
	ws := melody.New()

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		ws.HandleRequest(w, r)
	})

	ws.HandleMessage(func(s *melody.Session, b []byte) {
		w.input(b)
	})

	ticker := time.NewTicker(50 * time.Millisecond)
	go func() {
		for _ = range ticker.C {
			ws.Broadcast(w.frame())
		}
	}()
}
