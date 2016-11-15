package main

import (
	"fmt"
	"net/http"

	// TODO: absolute path -> relative path
	"github.com/mochi8k/ytan.ml/auth"
	"github.com/mochi8k/ytan.ml/chat"
	"github.com/mochi8k/ytan.ml/portal"
)

func main() {
	http.HandleFunc("/", portal.HTTPHandler)
	http.Handle("/chat", auth.Authenticate(chat.HTTPHandler))
	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/auth", auth.AuthHandler)

	fmt.Println("activated the web server on port 8080")
	http.ListenAndServe(":8080", nil)
}
