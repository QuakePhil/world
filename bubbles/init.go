package bubbles

import (
	"time"
)

var config struct {
	width, height float64
	fps           int
	frameDuration time.Duration
}

func init() {
	config.width = 300
	config.height = 150
	config.fps = 60
	config.frameDuration = time.Second / time.Duration(config.fps)
}
