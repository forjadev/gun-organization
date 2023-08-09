package handler

import (
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
	teamName := ctx.Param("teamname")

	if teamName == "" {
		sendError(ctx, 404, "Team name is required")
		return
	}

	members := []schemas.Member{}

	err := db.Joins("JOIN teams ON teams.team_id = member.team_id").Where("teams.name = ?", teamName).Find(&schemas.Member{})

	if err != nil {
		sendError(ctx, 404, "Team not found")
	}

	sendSuccess(ctx, "ListTeamMembers", members)
}
