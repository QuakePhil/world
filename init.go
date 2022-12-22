package main

var config struct {
	address string
	width   int
	height  int
}

func init() {
	config.address = "localhost:8081"
	config.width = 300
	config.height = 150
}
