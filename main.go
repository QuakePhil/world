package main

import (
	"log"
	"net/http"

	"github.com/olahol/melody"
)

func handleLocalFile(path, local string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, local)
	})
}

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

func main() {
	handleWebSockets("/ws")
	handleLocalFile("/", "client/world.html")
	handleLocalFile("/world.html", "client/world.html")
	handleLocalFile("/world.css", "client/world.css")

	address := "localhost:8081"
	log.Println("Listening for http and ws:", address)
	http.ListenAndServe(address, nil)
}
