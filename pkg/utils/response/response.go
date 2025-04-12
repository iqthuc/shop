package response

import (
	"encoding/json"
	"net/http"
	errs "shop/pkg/utils/errors"
	"shop/pkg/utils/messages"
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

func ErrorJson(w http.ResponseWriter, err errs.AppError, code int) {
	response := APIResponse{
		Message: err.Error(),
		Code:    code,
		Data:    nil,
		Status:  0,
	}
	JsonResponse(w, response)
}

func SuccessJson(w http.ResponseWriter, data any, message messages.AppMessage) {
	response := APIResponse{
		Message: string(message),
		Code:    http.StatusOK,
		Data:    data,
		Status:  1,
	}
	JsonResponse(w, response)
}
