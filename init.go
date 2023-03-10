package main

import (
	"time"
)

var config struct {
	address       string
	fps           int
	frameDuration time.Duration
}

func init() {
	config.address = "localhost:8081"
	config.fps = 60
	config.frameDuration = time.Second / time.Duration(config.fps)
}
