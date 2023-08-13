package repository

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _setupTestRouter() (route *gin.Engine) {
	route = gin.Default()
	route.GET("/readme", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return
}
