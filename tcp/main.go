package main

import (
	db "../database/client"
	_ "net/http/pprof"
	"./server"
	"log"
)

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Init client
	db.InitClient()

	// Start server
	server.Start()
}
