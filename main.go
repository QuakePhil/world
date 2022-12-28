package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"code/world/bouncyball"
	"code/world/cursor"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	handleWorld("/bouncyball/", "bouncyball")
	handleWorld("/cursor/", "cursor")

	log.Println("listening on", config.address)
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
				handleLocalFile(path, local+filename)
			}
			handleLocalFile(path+filename, local+filename)
		}
	}
}

func handleWorld(path, name string) {
	var w world
	switch name {
	case "bouncyball":
		w = new(bouncyball.World)
	case "cursor":
		w = new(cursor.World)
	default:
		panic("unknown world")
	}
	handleWebSockets(path+"ws", w)
	handleLocal(path, "client/")
	handleLocalFile(path+"client.js", name+"/client.js")
	log.Println("handling " + path)
}

func handleLocalFile(path, local string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, local)
	})
}
