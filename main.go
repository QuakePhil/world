package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"code/world/bouncyball"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var w world

func main() {
	if len(os.Args) < 2 {
		panic("Usage: go run . <world>")
	}
	switch os.Args[1] {
	case "bouncyball":
		w = new(bouncyball.World)
	default:
		panic(os.Args[1])
	}

	handleWebSockets("/ws", w)
	handleLocal("/", "client/")
	handleLocalFile("/client.js", "bouncyball/client.js")

	log.Println("Listening for http and ws on", config.address)
	http.ListenAndServe(config.address, nil)
}

func handleLocal(path, local string) {
	waitingForFirstHtml := true
	c, err := os.ReadDir(local)
	check(err)
	for _, entry := range c {
		if !entry.IsDir() {
			filename := entry.Name()
			if waitingForFirstHtml && filepath.Ext(filename) == ".html" {
				waitingForFirstHtml = false
				handleLocalFile("/", local+filename)
			}
			handleLocalFile("/"+filename, local+filename)
		}
	}
}

func handleLocalFile(path, local string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, local)
	})
}
