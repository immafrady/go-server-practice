package hero

import "net/http"

func AllHeroRoutes() {
	http.HandleFunc("/hero", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		switch request.Method {
		case http.MethodGet:
		case http.MethodPost:
		case http.MethodPut:
		case http.MethodDelete:
		}
	})
}
