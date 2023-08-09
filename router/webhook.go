package router

import (
	"github.com/forjadev/gun-organization/handler"
	"github.com/gin-gonic/gin"
)

func bindWebhookRoutes(group *gin.RouterGroup) {
	group.GET("/webhook", handler.GitHubWebhookHandler)
}
