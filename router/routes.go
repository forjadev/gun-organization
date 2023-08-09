package router

import (
	docs "github.com/forjadev/gun-organization/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

func InitializeRoutes() {
	r := gin.Default()
	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiGroup := r.Group(basePath)
	{
		bindActuatorsRoutes(apiGroup)
		bindWebhookRoutes(apiGroup)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := r.Run("0.0.0.0:" + port)
	if err != nil {
		panic(err)
	}
}
