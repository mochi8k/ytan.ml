package portal

import (
	"fmt"
	"html/template"
	"net/http"
)

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("portal")

	// TODO: absolute path -> relative path
	tpl, _ := template.ParseFiles("portal/templates/portal.tpl")
	tpl.Execute(w, "hello")
}
