package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

// Initialize sets up the router and starts the server
func Initialize() {
	// Create a new Gin router with default middleware
	r := gin.Default()

	// Initialize the routes for the application
	initializeRoutes(r)

	// Get the port number from the environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server on the specified port
	r.Run("0.0.0.0:" + port)
}
