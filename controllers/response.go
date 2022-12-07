package controllers

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(code int, message string, data interface{}) response {
	if data == nil {
		data = map[string]any{}
	}

	return response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
