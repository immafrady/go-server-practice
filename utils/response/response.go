package response

import "encoding/json"

type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

func newResponse(code int, data interface{}, msg string) *response {
	return &response{Code: code, Data: data, Msg: msg}
}

func (r *response) ToJson() []byte {
	js, err := json.Marshal(r)
	if err != nil {
		resp := JSONConvertError(err)
		return resp.ToJson()
	}
	return js
}
