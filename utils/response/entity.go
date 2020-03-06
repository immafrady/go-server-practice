package response

func SuccessResponse(data interface{}, msg string) *response {
	return newResponse(1000, data, msg)
}

func JSONConvertError(err error) *response {
	return newResponse(1001, nil, err.Error())
}

func DbError(err error) *response {
	return newResponse(1002, nil, err.Error())
}

func ParamError(err error) *response {
	return newResponse(1003, nil, err.Error())
}

func NotSupportedMethodError(err error) *response {
	return newResponse(1004, nil, err.Error())
}

func UnknownError(err error) *response {
	return newResponse(9999, nil, err.Error())
}
