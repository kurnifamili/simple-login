package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../database"
	"runtime"
)

var (
	dbClient = database.NewDatabaseClient()
)

type LoginRequest struct {
	Username string
	Password string
}

func main() {
	runtime.GOMAXPROCS(30)
	http.HandleFunc("/login", LoginHandler)
	http.ListenAndServe(":8090", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest

	// Parse request body into User object
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, _, _, err := dbClient.GetUser(loginReq.Username, loginReq.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "UserID:", userID)
	fmt.Fprintln(w, "Login successful!")
}
