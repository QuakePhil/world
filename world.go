package main

import (
	"net/http"
	"time"

	"github.com/olahol/melody"
)

type world interface {
	input([]byte)
	frame() []byte
}

func handleWebSockets(path string, w world) {
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
