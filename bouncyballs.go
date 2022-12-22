package main

import (
	"log"
)

type bouncyball struct {
	x, y float64
}

var world []bouncyball

func bouncyballs() {
	ws.Broadcast([]byte("10 10 20 20"))
}
