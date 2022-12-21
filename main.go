package main

import (
	"log"
	"net/http"

	"github.com/olahol/melody"
)

var m = melody.New()

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client/world.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	m.HandleMessage(func(s *melody.Session, b []byte) {
		log.Println(string(b))
		m.Broadcast(b)
	})

	http.ListenAndServe("localhost:8080", nil)
}
