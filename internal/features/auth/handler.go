package auth

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	errs "shop/pkg/utils/errors"
	"shop/pkg/utils/messages"
	"shop/pkg/utils/response"
)

type UseCase interface {
	SignUp(ctx context.Context, input signUpInput) (*signUpResult, error)
}

type handler struct {
	useCase UseCase
}

func NewHandler(useCase UseCase) handler {
	return handler{
		useCase: useCase,
	}
}

func (h handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req signUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	input := signUpInput{
		Email:    req.Email,
		Password: req.Password,
	}
	result, err := h.useCase.SignUp(r.Context(), input)
	if err != nil {
		switch {
		case errors.Is(err, errs.VaidationFailed):
			response.ErrorJson(w, errs.PrettyValidationErrors(err), http.StatusBadRequest)
		case errors.Is(err, errs.EmailAlready):
			response.ErrorJson(w, errs.EmailAlready, http.StatusBadRequest)
		default:
			response.ErrorJson(w, errs.SomethingWrong, http.StatusInternalServerError)
		}
		slog.Warn("sign up failed", slog.String("error", err.Error()))
		return
	}

	resp := signUpResponse{
		Email:     result.email,
		CreatedAt: result.createdAt,
	}
	response.SuccessJson(w, resp, messages.SignUpSuccess)
	slog.Info("sign up success")
}
