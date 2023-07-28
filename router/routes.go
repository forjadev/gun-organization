package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	basePath := "/api/v1"

	v1 := r.Group(basePath)

	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
				"status":  http.StatusOK,
			})
		})
	}
}
