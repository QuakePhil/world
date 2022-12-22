package main

import (
	"net/http"
	"time"

	"github.com/olahol/melody"
)

type world interface {
	Input([]byte)
	Frame() []byte
}

func handleWebSockets(path string, w world) {
	ws := melody.New()

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		ws.HandleRequest(w, r)
	})

	ws.HandleMessage(func(s *melody.Session, b []byte) {
		w.Input(b)
	})

	ticker := time.NewTicker(config.frameDuration)
	go func() {
		for _ = range ticker.C {
			ws.Broadcast(w.Frame())
		}
	}()
}
