package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../common"
	"../tcp"
)

var (
	tcpClient tcp.ITcpClient
)

const (
	HttpPort = "8090"
)

func main() {
	//runtime.GOMAXPROCS(20)

	// init tcpClient
	tcpClient = tcp.NewTcpClient()

	// start server server
	http.HandleFunc("/login", LoginHandler)
	fmt.Printf("Starting HTPP server at port %s..\n", HttpPort)
	http.ListenAndServe(":"+HttpPort, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq common.LoginRequest

	// Parse request body into User object
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := tcpClient.SendLoginRequest(&loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "UserID:", userID)
	fmt.Fprintln(w, "Login successful!")
}
