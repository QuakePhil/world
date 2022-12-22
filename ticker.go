package main

import (
	"time"
)

func ticker() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for _ = range ticker.C {
			bouncyballs()
		}
	}()
}
