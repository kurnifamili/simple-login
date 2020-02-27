package main

import (
	tcp "../tcp/client"
	_ "net/http/pprof"
	"./server"
	"log"
)

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// init clients
	tcp.InitClient()

	// start HTTP server
	server.Start()
}
