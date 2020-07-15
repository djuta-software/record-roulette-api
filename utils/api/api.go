package api

import (
	"encoding/json"
	"net/http"
)

type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// SendOkResponse returns a 200 response
func SendOkResponse(writer http.ResponseWriter, value interface{}) {
	sendResponse(writer, http.StatusOK, value)
}

// SendBadRequestResponse returns a 400 response
func SendBadRequestResponse(writer http.ResponseWriter, message string) {
	status := http.StatusBadRequest
	value := apiError{Status: status, Message: message}
	sendResponse(writer, status, value)
}

// SendInternalServerError returns a 500 response
func SendInternalServerError(writer http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	message := "An unexpected error occured"
	value := apiError{Status: status, Message: message}
	sendResponse(writer, status, value)
}

func sendResponse(writer http.ResponseWriter, statusCode int, value interface{}) {
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(value)
}
