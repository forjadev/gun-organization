package handler

import (
	"net/http"

	"github.com/forjadev/gun-organization/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Get organization teams
// @Description Get all teams from forja organization
// @Tags Teams
// @Accept json
// @Produce json
// @Success 200 {object} TeamsServerResponse
// @Router /teams [get]
func GetTeamsHandler(ctx *gin.Context) {
	teams := schemas.TeamsResponse{
		Message: "teams",
		Status: http.StatusOK,
	}
	sendSuccess(ctx, "GetTeamsHandler", teams)
}