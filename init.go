package main

var config struct {
	address string
	width   float64
	height  float64
}

func init() {
	config.address = "localhost:8081"
	config.width = 300
	config.height = 150
}
