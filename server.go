package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my page")
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("templates/chat.tpl")
	tpl.Execute(w, "hello")
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/chat", chatHandler)

	fmt.Println("activated the web server on port 8080")
	http.ListenAndServe(":8080", nil)

}
