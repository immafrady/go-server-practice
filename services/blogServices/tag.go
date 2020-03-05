package blogServices

import (
	"fradyspace.com/go-server-practice/models/dao"
	"fradyspace.com/go-server-practice/utils/response"
	"net/http"
)

func BlogTagPostService(tagName string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		err := dao.NewTag(tagName)
		if err != nil {
			w.Write(response.DbError(err).ToJson())
		} else {
			w.Write(response.SuccessResponse(nil, "成功添加").ToJson())
		}
	}
}

func BlogTagDeleteService(tagName string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		err := dao.DeleteTag(tagName)
		if err != nil {
			w.Write(response.DbError(err).ToJson())
		} else {
			w.Write(response.SuccessResponse(nil, "成功删除").ToJson())
		}
	}
}
