package server

import "net/http"

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewOKResponse(data any) *response {
	return &response{
		Status:  http.StatusOK,
		Message: "",
		Data:    data,
	}
}

func NewErrorResponse(status int, message string) *response {
	return &response{
		Status:  status,
		Message: message,
	}
}
