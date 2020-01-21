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

	tcpReq := constructTcpLoginRequest(loginReq)
	client := tcpClient.GetClient()
	userID, err := client.SendRequest(tcpReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "UserID:", userID)
	fmt.Fprintln(w, "Login successful!")
}

func constructTcpLoginRequest(request common.LoginRequest) string {
	return common.TcpLoginRequest.ToString() + " " + request.Username + " " + request.Password + common.TcpMsgDelimiterStr
}