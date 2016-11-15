package auth

import (
	"log"
	"net/http"
)

type basicHttpHandler func(http.ResponseWriter, *http.Request)

type authHandler struct {
  nextHandler basicHttpHandler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Location", "/login")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func Authenticate(handler basicHttpHandler) http.Handler {
  log.Println("Authenticate")
  return &authHandler{nextHandler: handler}
}
