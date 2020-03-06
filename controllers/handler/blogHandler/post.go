package blogHandler

import (
	"fradyspace.com/go-server-practice/utils/response"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

const PrefixPost = "/blog/post/"

var prefixLenPost = len(PrefixPost)

func PostRoutes() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tagName := req.URL.Path[prefixLenPost:]
		switch tagName {
		case "list":
			if req.Method == http.MethodGet {

			} else {
				w.Write(response.NotSupportedMethodError(errors.New("不支持的请求方法：" + req.Method)).ToJson())
			}
		case "new":
			if req.Method == http.MethodPost {

			} else {
				w.Write(response.NotSupportedMethodError(errors.New("不支持的请求方法：" + req.Method)).ToJson())
			}
		default:
			id, err := strconv.Atoi(tagName)
			if err != nil {
				w.Write(response.ParamError(err).ToJson())
				return
			}
			switch req.Method {
			case http.MethodPut:
			case http.MethodDelete:

			}
		}
	}
}
