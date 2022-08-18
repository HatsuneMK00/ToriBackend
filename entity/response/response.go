package response

import "net/http"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(message string, data interface{}) Response {
	return Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}
}
