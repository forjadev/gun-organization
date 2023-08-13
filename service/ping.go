package service

import (
	"github.com/forjadev/gun-organization/handler"
	"net/http"

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
	//pingData := schemas.PingResponse{
	//	Message: "pong",
	//	Status:  http.StatusOK,
	//}

	// TODO: fix error handling
	handler.SendError(ctx, http.StatusBadRequest, "ok")
}
