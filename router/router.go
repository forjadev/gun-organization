package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// Initialize sets up the router and starts the server
func Initialize() error {
	// Create a new Gin router with default middleware
	r := gin.Default()

	// Initialize the routes for the application
	initializeRoutes(r)

	// Get the port number from the environment variables

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = fmt.Sprintf(":%s", port)

	// Start the server on the specified port
	if err := r.Run("localhost" + port); err != nil {
		return fmt.Errorf("set route to run: %v", err)
	}

	return nil
}
