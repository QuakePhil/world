package main

import (
	"log"
	"net/http"

	"github.com/olahol/melody"
)

var m = melody.New()

func handleLocalFile(path, local string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, local)
	})
}

func handleWebSockets(path string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	m.HandleMessage(func(s *melody.Session, b []byte) {
		log.Println(string(b))
		m.Broadcast(b)
	})
}

func main() {
	handleWebSockets("/ws")
	handleLocalFile("/", "client/world.html")
	handleLocalFile("/world.html", "client/world.html")
	handleLocalFile("/world.css", "client/world.css")

	http.ListenAndServe("localhost:8081", nil)
}
