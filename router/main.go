package router

import (
	"fradyspace.com/go-server-practice/app/controller/hero"
	"fradyspace.com/go-server-practice/app/controller/statistic"
	"net/http"
)

func Routers() {
	http.HandleFunc("/statistic", statistic.MainController)
	http.HandleFunc("/hero", hero.MainController)
}
