package server

import (
	"net/http"
	"fmt"

	"../handler"
	"../../common"
)

func Start() {
	http.HandleFunc("/login", handler.LoginHandler)
	fmt.Printf("Starting HTPP server at port %s..\n", common.HttpPort)
	http.ListenAndServe(":"+common.HttpPort, nil)
}