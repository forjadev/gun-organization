package router

import (
	docs "github.com/forjadev/gun-organization/docs"
	"github.com/forjadev/gun-organization/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// initializeRoutes sets up the routes for the application
func initializeRoutes(r *gin.Engine) {
	// Set the base path for all routes
	basePath := "/api/v1"

	// Set the base path for Swagger documentation
	docs.SwaggerInfo.BasePath = basePath

	// Create a new router group for version 1 of the API
	v1 := r.Group(basePath)
	//TODO:team := service.NewTeamService()
	webhook := service.NewWebhookService()
	// Define the routes for version 1 of the API
	{
		// Define a GET route for the /ping endpoint
		v1.GET("/ping", service.PingServerHandler)
		//TODO:v1.GET("/teams", team.ListTeams)
		v1.POST("/webhook", webhook.GetWebhook)
	}

	// Initialize Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
