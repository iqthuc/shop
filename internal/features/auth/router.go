package auth

import (
	"database/sql"
	"log"
	"net/http"
)

func SetupRouter(handler handler) http.Handler {
	log.Println("setup router")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /sign_up", handler.SignUp)
	// mux.HandleFunc("GET /logOut", handler.LogOut)
	return mux
}

func RegisterRoute(mainRoute *http.ServeMux, db *sql.DB) {
	authRepo := NewRepository(db)
	authUseCase := NewUsecase(authRepo)
	authHandler := NewHandler(authUseCase)
	authRoute := SetupRouter(authHandler)
	mainRoute.Handle("/auth/", http.StripPrefix("/auth", authRoute))
}
