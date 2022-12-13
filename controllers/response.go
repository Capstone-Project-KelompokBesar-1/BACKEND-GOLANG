package controllers

type responseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type responseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(code int, message string, data interface{}) any {
	if code >= 400 {
		return responseError{
			Code:    code,
			Message: message,
		}
	}

	if data == nil {
		data = map[string]any{}
	}

	return responseSuccess{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
