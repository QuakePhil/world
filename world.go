package main

import (
	"bytes"
	"fmt"
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
		bouncyballSpawn(b)
	})

	ticker := time.NewTicker(50 * time.Millisecond)
	go func() {
		for _ = range ticker.C {
			ws.Broadcast(bouncyballs())
		}
	}()
}

var world []bouncyball

func bouncyballs() []byte {
	var b bytes.Buffer
	for i := range world {
		// fmt.Println("broadcast:", world[i])
		b.Write(world[i].bytes())
		b.Write([]byte(" "))
		world[i].think()
	}
	return b.Bytes()
}

func bouncyballSpawn(b []byte) {
	world = append(world, bouncyballFromBytes(b))
	fmt.Println("spawned:", world[len(world)-1])
}
