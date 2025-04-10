package auth

import (
	"log"
	"net/http"
)

func SetupRouter(handler handler) http.Handler {
	log.Println("setup router")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /login", handler.Login)
	// mux.HandleFunc("GET /logOut", handler.LogOut)
	return mux
}

func RegisterRoute(mainRoute *http.ServeMux) {
	authRepo := NewRepository()
	authUseCase := NewUsecase(authRepo)
	authHandler := NewHandler(authUseCase)
	authRoute := SetupRouter(authHandler)
	mainRoute.Handle("/auth/", http.StripPrefix("/auth", authRoute))
}
