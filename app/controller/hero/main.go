package hero

import (
	"net/http"
)

func MainController(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getHeroes(w, req)
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	}
}

func getHeroes(w http.ResponseWriter, req *http.Request) {

}
