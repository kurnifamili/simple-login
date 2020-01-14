package main

import (
	tcp "../tcp/client"
	"./server"
)

func main() {
	// init clients
	tcp.InitClient()

	// start HTTP server
	server.Start()
}
