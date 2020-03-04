package controllers

import (
	"fradyspace.com/go-server-practice/controllers/handler"
	"fradyspace.com/go-server-practice/middleware"
	"fradyspace.com/go-server-practice/middleware/request"
)

func StartRouting() *middleware.Router {
	router := middleware.NewRouter()
	router.Use(request.Logger)
	router.Add("/hero", handler.HeroRoutes())
	router.Add("/page", handler.PageRoutes())

	return router
}
