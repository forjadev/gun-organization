package handler

import (
	"net/http"

	"github.com/forjadev/gun-organization/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Ping the server
// @Description Ping the server to check if it is running
// @Tags Ping
// @Accept json
// @Produce json
// @Success 200 {object} PingServerResponse
// @Router /ping [get]
func PingServerHandler(ctx *gin.Context) {
	pingData := schemas.PingResponse{
		Message: "pong",
		Status:  http.StatusOK,
	}

	sendSuccess(ctx, "PingServerHandler", pingData)
}
