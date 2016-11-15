package main

import (
	"fmt"
	"net/http"

	// TODO: absolute path -> relative path
	"github.com/mochi8k/ytan.ml/chat"
	"github.com/mochi8k/ytan.ml/portal"
	"github.com/mochi8k/ytan.ml/auth"
)

func main() {
	http.HandleFunc("/", portal.HTTPHandler)
	http.Handle("/chat", auth.Authenticate(chat.HTTPHandler))
	http.HandleFunc("/login", auth.HTTPHandler)

	fmt.Println("activated the web server on port 8080")
	http.ListenAndServe(":8080", nil)
}
