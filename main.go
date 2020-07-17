package main

import(
	"fmt"
	"kafkabrowser/server"
)

const (
	version = "v0.0.1"
)

func SpalshScreen() {
	fmt.Printf("Starting Kafka Browser\n")
	fmt.Printf("Version: %s",version)
	fmt.Printf("---------------------------------------------------------------------\n")
}

func main() {
	server.ServeApp()
}