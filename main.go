package main

import (
	"github.com/Miguelluiddev/Gopportunities/config"
	"github.com/Miguelluiddev/Gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.ErrorF("config initialization error: %v", err)

	}

	router.Initialize()
}
