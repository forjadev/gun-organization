package main

import (
	"log"

	"github.com/forjadev/gun-organization/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("could not load environment variables: %v", err)
	}

	//Initialize configs
	//err = config.Init()

	if err != nil {
		log.Fatalf("could not initialize config: %v", err)
	}
	// Initialize Router
	router.InitializeRoutes()
}
