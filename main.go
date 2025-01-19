package main

import (
	"github.com/CastroAproved/apiGO/config"
	"github.com/CastroAproved/apiGO/router"
)

var (
	logger *config.Logger
)


func main() {

logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
		//panic(err)
	}

// Initialize the router
	router.Initialize()

}