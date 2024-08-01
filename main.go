package main

import (
	"fmt"
	"net/http"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Version: 2.0.0")
}

func main() {
	http.HandleFunc("/", versionHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
