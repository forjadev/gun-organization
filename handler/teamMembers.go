package handler

import (
	"net/http"

	"github.com/forjadev/gun-organization/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Returns the list of team members
// @Description Returns the list of members of specific team
// @Tags Teams
// @Accept json
// @Produce json
// @Success 200 {object} ListTeamMembersResponse
// @Router /team/{teamname}/members [get]
func ListTeamMembersHandler(ctx *gin.Context) {
	// DB not initialized yet!
	var db

	teamName := ctx.Param("teamname")

	if teamName == "" {
		sendError(ctx, 404, "Team name is required")
		return
	}

	if result := db.Joins("JOIN teams ON teams.team_id = member.team_id").Where("teams.name = ?", teamName).Find(&schema.Member{}); result.Error != nil {
		sendError(ctx, 404, "Team not found")
	}

	sendSuccess(ctx, "ListTeamMembers", result)
}
