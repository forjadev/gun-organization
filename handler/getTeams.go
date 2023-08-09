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
	teams := []schemas.Team{}

	if err := db.Find(&teams).Error;err!=nil{
		sendError(ctx, http.StatusInternalServerError, "error listing teams")
		return
	}

	sendSuccess(ctx, "GetTeamsHandler", teams)
}