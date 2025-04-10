package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type UserCase interface {
	Login(ctx context.Context, email string, password string) error
}

type handler struct {
	useCase UserCase
}

func NewHandler(useCase UserCase) handler {
	return handler{
		useCase: useCase,
	}
}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("login")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Hello")
}
