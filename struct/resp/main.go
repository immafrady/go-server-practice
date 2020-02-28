package resp

import (
	"encoding/json"
	"encoding/xml"
)

type JsonResponse struct {
	Code int         `json:"code",xml:""`
	Data interface{} `json:"code",xml:""`
	Msg  string      `json:"msg",xml:""`
}

func NewFatalJsonResponse(msg string) *JsonResponse {
	return &JsonResponse{Code: FAIL, Msg: msg}
}

func (j *JsonResponse) SetCode(code int) {
	j.Code = code
}

func (j *JsonResponse) SetData(data interface{}) {
	j.Data = data
}

func (j *JsonResponse) SetMsg(msg string) {
	j.Msg = msg
}

func (j *JsonResponse) toJsonString() []byte {
	bytes, err := json.Marshal(j)
	if err != nil {
		logger.Errorf(err.Error())
	}
	return bytes
}

func (j *JsonResponse) toXMLString() []byte {
	bytes, err := xml.Marshal(j)
	if err != nil {
		logger.Errorf(err.Error())
	}
	return bytes
}
