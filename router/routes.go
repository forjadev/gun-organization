package router

import (
	docs "github.com/forjadev/gun-organization/docs"
	"github.com/forjadev/gun-organization/handler"
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

	// Define the routes for version 1 of the API
	{
		// Define a GET route for the /ping endpoint
		v1.GET("/ping", handler.PingServerHandler)
		v1.GET("/teams/:teamname/members", handler.ListTeamMembersHandler)
	}

	// Initialize Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
