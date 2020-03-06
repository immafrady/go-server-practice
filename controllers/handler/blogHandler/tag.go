package blogHandler

import (
	"fradyspace.com/go-server-practice/services/blogServices"
	"net/http"
)

const PrefixTag = "/blog/tag/"

var prefixLenTag = len(PrefixTag)

func TagRoutes() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tagName := req.URL.Path[prefixLenTag:]
		switch req.Method {
		case http.MethodGet:
			blogServices.BlogTagGetService(tagName)(w, req)
		case http.MethodPost:
			// 新增标签
			blogServices.BlogTagPostService(tagName)(w, req)
		case http.MethodDelete:
			// 删除标签
			blogServices.BlogTagDeleteService(tagName)(w, req)
		}
	}
}
