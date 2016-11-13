package chat

import (
	"html/template"
	"log"
	"net/http"
)

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("start chat")

	data := map[string]interface{}{
		"Host": r.Host,
	}

	// TODO: absolute path -> relative path
	// TODO: コンパイルを初回のみにする. sync.Once
	tpl, _ := template.ParseFiles("chat/templates/chat.tpl")
	tpl.Execute(w, data)
}
