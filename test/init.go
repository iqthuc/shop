package test

import (
	"shop/internal/app"
	"shop/internal/infrastructure/cache"
	"shop/internal/infrastructure/database/store"
	"shop/internal/infrastructure/server"
	"shop/pkg/token"

	"github.com/go-playground/validator/v10"
)

var testApp *app.Application

func init() {
	testConfig := app.LoadConfig("../configs", "config.dev", "yaml")

	// logger.ConfigureLogger(testConfig.Logger)
	server := server.New(testConfig.Server)
	store := store.NewPostgresStore(testConfig.Database)
	validator := validator.New()
	tkMaker := token.NewJwtMaker(testConfig.Token.SecretKey)
	redis := cache.NewRedisClient(testConfig.Redis)
	testApp = app.NewApp(server, store, validator, tkMaker, redis)
	testApp.RegisterRoutes()
}
