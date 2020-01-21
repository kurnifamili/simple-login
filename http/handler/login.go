package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	client := tcpClient.GetClient()
	tcpReq := constructTcpLoginRequest(loginReq)
	userID, err := client.SendRequest(tcpReq)

	// time1 := time.Now()
	// userID, err := client.SendRequestBytes(constructTcpLoginRequestBytes(loginReq))
	// time2 := time.Now()
	// log.Printf("HEREE request time %s", time2.Sub(time1))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "UserID:", userID)
	fmt.Fprintln(w, "Login successful!")
}

func constructTcpLoginRequest(request common.LoginRequest) string {
	return common.TcpLoginRequest.ToString() + common.TcpMsgSeparator + request.Username + common.TcpMsgSeparator + request.Password + common.TcpMsgDelimiterStr
}

// func constructTcpLoginRequestBytes(request common.LoginRequest) []byte {
// 	arr := append(common.TcpLoginRequest.ToBytes(), []byte(common.TcpMsgSeparator+request.Username+common.TcpMsgSeparator+request.Password+common.TcpMsgDelimiterStr)...)
// 	return arr
// }
