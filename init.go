package main

import (
	"time"
)

var config struct {
	address       string
	width         float64
	height        float64
	fps           int
	frameDuration time.Duration
}

func init() {
	config.address = "localhost:8081"
	config.width = 300
	config.height = 150
	config.fps = 60
	config.frameDuration = time.Second / time.Duration(config.fps)
}
