package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"code/world/bouncyball"
	"code/world/bubbles"
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
	case "bubbles":
		w = new(bubbles.World)
	default:
		panic("unknown world")
	}

	handleWebSockets("/ws", w)
	handleLocal("/", "client/")
	handleLocalFile("/client.js", os.Args[1]+"/client.js")

	log.Println(os.Args[1], "listening on", config.address)
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
