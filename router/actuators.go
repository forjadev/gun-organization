package router

import (
	"github.com/forjadev/gun-organization/handler"
	"github.com/gin-gonic/gin"
)

func bindActuatorsRoutes(group *gin.RouterGroup) {
	group.GET("/ping", handler.PingServerHandler)
}
