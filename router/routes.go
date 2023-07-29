package router

import (
	"net/http"

	docs "github.com/forjadev/gun-organization/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := r.Group(basePath)

	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
				"status":  http.StatusOK,
			})
		})
	}

	// Initialize swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
