package handler

import (
	"fradyspace.com/go-server-practice/services"
	"fradyspace.com/go-server-practice/utils/logger"
	"net/http"
)

func HeroRoutes() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		switch request.Method {
		case http.MethodGet:
			logger.CheckError("HeroGetService", services.HeroGetService(writer, request))
		case http.MethodPost:
			logger.CheckError("HeroPostService", services.HeroPostService(writer, request))
		case http.MethodPut:
			logger.CheckError("HeroPutService", services.HeroPutService(writer, request))
		case http.MethodDelete:
			logger.CheckError("HeroDeleteService", services.HeroDeleteService(writer, request))
		}
	}
}
