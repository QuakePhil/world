package main

import (
	"net/http"
	"time"

	"github.com/olahol/melody"
)

var ws *melody.Melody

func handleWebSockets(path string) {
	ws = melody.New()

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		ws.HandleRequest(w, r)
	})

	ws.HandleMessage(func(s *melody.Session, b []byte) {
		bouncyball_spawn(b)
	})

	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for _ = range ticker.C {
			bouncyballs()
		}
	}()
}
