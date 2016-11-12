package chat

import (
	"fmt"
	"html/template"
	"net/http"
)

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("chat")

	// TODO: absolute path -> relative path
	tpl, _ := template.ParseFiles("chat/templates/chat.tpl")
	tpl.Execute(w, "hello")
}
