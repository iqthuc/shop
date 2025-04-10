package app

import (
	"fmt"
	"shop/internal/infrastructure/config"
)

func Bootstrap() {
	appConfig, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}

	app := NewApp(appConfig)
	app.Run()
}
