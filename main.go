package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", LoginHandler)
	http.ListenAndServe(":8090", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login successful!")
}
