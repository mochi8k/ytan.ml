package auth

import (
	"html/template"
	"log"
	"net/http"
)

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/login")
	tpl, _ := template.ParseFiles("auth/templates/login.tpl")
	tpl.Execute(w, "")
}
