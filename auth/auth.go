package auth

import (
	"log"
	"net/http"
)

type basicHTTPHandler func(http.ResponseWriter, *http.Request)

type authRouter struct {
	nextHandler basicHTTPHandler
}

func (ar *authRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth")

	if err == http.ErrNoCookie || cookie.Value == "" {

		log.Println("未認証")

		// 未認証
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)

	} else if err != nil {

		log.Println("panic")

		// 想定外のエラー
		panic(err.Error())

	} else {

		log.Println("認証済み")

		// 認証済み
		ar.nextHandler(w, r)

	}

}

func Authenticate(handler basicHTTPHandler) http.Handler {
	return &authRouter{nextHandler: handler}
}
