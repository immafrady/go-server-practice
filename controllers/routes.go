package controllers

import (
	"fradyspace.com/go-server-practice/controllers/handler/blogHandler"
	"fradyspace.com/go-server-practice/middleware"
	"fradyspace.com/go-server-practice/middleware/request"
)

func StartRouting() *middleware.Router {
	router := middleware.NewRouter()
	router.Use(request.Logger)
	router.Add(blogHandler.PrefixTag, blogHandler.TagRoutes())

	return router
}
