package auth

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/login")
	tpl, _ := template.ParseFiles("auth/templates/login.tpl")
	tpl.Execute(w, "")
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/auth")

	// TODO: セッション管理
	r.ParseForm()
	log.Println(r.Form)

	// 有効期限: 1分
	expiration := time.Now().Add(1 * time.Minute)

	http.SetCookie(w, &http.Cookie{
		Name:    "auth",
		Value:   "session_id", // TODO ユニーク
		Expires: expiration,
	})

	w.Header().Set("Location", "/portal")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
