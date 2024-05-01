package http

import "net/http"

func NewSuccessCreatedResponse() DefaultResponse {
	return DefaultResponse{
		Code:    http.StatusCreated,
		Status:  "Success",
		Message: "Success Created Data",
	}
}

func NewSuccessDefaultResponse() DefaultResponse {
	return DefaultResponse{
		Code:   http.StatusOK,
		Status: "Success",
	}
}
