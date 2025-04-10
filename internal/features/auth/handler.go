package auth

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"shop/pkg/utils"
)

type handler struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) handler {
	return handler{
		useCase: useCase,
	}
}

type UseCase interface {
	SignUp(ctx context.Context, input signUpInput) (signUpResult, error)
}

func (h handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req signUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		utils.ErrorJsonResponse(w, http.StatusBadRequest, utils.ErrorMessages.InvalidRequest)
		return
	}

	input := signUpInput{
		req.Username,
		req.Email,
		req.Password,
	}

	signUpResponse, err := h.useCase.SignUp(r.Context(), input)
	if err != nil {
		log.Println(err)
		utils.ErrorJsonResponse(w, http.StatusInternalServerError, utils.ErrorMessages.SomethingErrors)
		return
	}

	utils.SuccessJsonResponse(w, signUpResponse, "signup success")
}
