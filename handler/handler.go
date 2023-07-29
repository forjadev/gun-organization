package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Ping the server
// @Description Ping the server to check if it is running
// @Tags Ping
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Returns a JSON object with the message and status code"
// @Router /ping [get]
func PingServerHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"status":  http.StatusOK,
	})
}
