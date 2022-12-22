package main

var config struct {
	address string
}

func init() {
	config.address = "localhost:8081"
}
