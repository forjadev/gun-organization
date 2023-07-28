package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()

	initializeRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run("0.0.0.0:" + port)
}
