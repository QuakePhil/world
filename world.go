package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/olahol/melody"
)

func handleWebSockets(path string) {
	ws := melody.New()

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		ws.HandleRequest(w, r)
	})

	ws.HandleMessage(func(s *melody.Session, b []byte) {
		input(b)
	})

	ticker := time.NewTicker(50 * time.Millisecond)
	go func() {
		for _ = range ticker.C {
			ws.Broadcast(frame())
		}
	}()
}

var world []bouncyball

func frame() []byte {
	var b bytes.Buffer
	for i := range world {
		b.Write(world[i].bytes())
		b.Write([]byte(" "))
		world[i].think()
	}
	return b.Bytes()
}

func input(b []byte) {
	world = append(world, bouncyballFromBytes(b))
	log.Println("spawned:", world[len(world)-1])
}
