package main

import (
	"log"
	"net/http"
)

func main() {
	handleWebSockets("/ws") // websocket.go
	handleLocalFile("/", "client/world.html")
	handleLocalFile("/world.html", "client/world.html")
	handleLocalFile("/world.css", "client/world.css")
	handleLocalFile("/js/bouncyballs.js", "client/js/bouncyballs.js")
	listen("localhost:8081")
}

func handleLocalFile(path, local string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, local)
	})
}

func listen(address string) {
	log.Println("Listening for http and ws on", address)
	http.ListenAndServe(address, nil)
}
