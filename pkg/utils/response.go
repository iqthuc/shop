package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Status  int    `json:"status"`
}

func JsonResponse(w http.ResponseWriter, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

func ErrorJsonResponse(w http.ResponseWriter, code int, message string) {
	response := APIResponse{
		Message: message,
		Code:    code,
		Data:    nil,
		Status:  0,
	}
	JsonResponse(w, response)
}

func SuccessJsonResponse(w http.ResponseWriter, data any, message string) {
	response := APIResponse{
		Message: message,
		Code:    http.StatusOK,
		Data:    data,
		Status:  1,
	}
	JsonResponse(w, response)
}
