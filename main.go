package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", TestServer)
	http.ListenAndServe(":8080", nil)
}

func TestServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TEST! Subpath %s", r.URL.Path[1:])
}
