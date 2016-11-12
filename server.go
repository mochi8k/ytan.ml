package main

import (
	"fmt"
	"net/http"

	"./chat"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my page")
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/chat", chat.HTTPHandler)

	fmt.Println("activated the web server on port 8080")
	http.ListenAndServe(":8080", nil)
}
