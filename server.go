package main

import (
	"fmt"
	"net/http"

	"./chat"
	"./portal"
)

func main() {
	http.HandleFunc("/", portal.HTTPHandler)
	http.HandleFunc("/chat", chat.HTTPHandler)

	fmt.Println("activated the web server on port 8080")
	http.ListenAndServe(":8080", nil)
}
