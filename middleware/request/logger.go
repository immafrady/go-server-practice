package request

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//defer func() {
		//	if err := recover(); err != nil {
		//		log.Printf("Lethal Error! [Path: %v, Method: %v] - %v\n", request.URL.Path, request.Method, err)
		//	}
		//}()
		log.Printf("[Origin:%v](Method:%v): %v\n", request.Header.Get("Origin"), request.Method, request.URL.Path)
		next.ServeHTTP(writer, request)
	})
}
