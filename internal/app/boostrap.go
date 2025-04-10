package app

import (
	"fmt"
	"net/http"
	"shop/internal/config"
)

func Bootstrap() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}

	rootRouter := http.NewServeMux()
	app := Application{
		config: cfg,
		router: rootRouter,
	}

	fmt.Println("bootstrap")
	app.Run()
}
