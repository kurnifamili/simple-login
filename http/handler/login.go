package handler

import (
	"encoding/json"
	"net/http"
	"fmt"

	"../../common"
	tcpClient "../../tcp/client"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq common.LoginRequest

	// Parse request body into User object
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := tcpClient.GetClient().SendLoginRequest(&loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "UserID:", userID)
	fmt.Fprintln(w, "Login successful!")
}