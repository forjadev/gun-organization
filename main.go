package main

import (
	"github.com/forjadev/gun-organization/repository"
	"log"

	"github.com/forjadev/gun-organization/router"
	"github.com/joho/godotenv"
)

// Load environment variables, initialize config and router
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("could not load environment variables: %v", err)
	}

	db := repository.NewDatabase()
	if err := db.Connect(); err != nil {
		log.Fatalf("could not initialize database: %v", err)
	}

	if err := router.Initialize(); err != nil {
		log.Fatalf("could not initialize router: %v", err)
	}
}
