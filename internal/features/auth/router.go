package auth

import (
	"net/http"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/routable"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type module struct {
	store     store.Store
	validator *validator.Validate
}

func NewModule(s store.Store, v *validator.Validate) routable.Routable {
	return &module{
		store:     s,
		validator: v,
	}
}
func (m *module) InitRoutes() http.Handler {
	repo := NewRepository(m.store)
	useCase := NewUsecase(repo, *m.validator)
	handler := NewHandler(useCase)

	mux := chi.NewMux()
	mux.Get("/sign-up", handler.SignUp)
	// mux.HandleFunc("GET /login", handler.Login)
	return mux
}
