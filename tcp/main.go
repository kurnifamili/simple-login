package main

import (
	db "../database/client"
	"./server"
)

func main() {
	// Init client
	db.InitClient()

	// Start server
	server.Start()
}
