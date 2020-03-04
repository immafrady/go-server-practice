package handler

import (
	"html/template"
	"net/http"
)

var t template.Template

func PageRoutes() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		http.NotFound(writer, request)
		//http.StatusFound
		//t = template.Must(template.ParseFiles("./view.html"))
		//t.Execute(writer, nil)
	}
}
