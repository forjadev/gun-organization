package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize router
	r := gin.Default()

	// Initialize routes
	initializeRoutes(r)

	// Get the port from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run the server
	r.Run("0.0.0.0:" + port)
}
