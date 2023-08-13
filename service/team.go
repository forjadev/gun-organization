package service

import (
	"github.com/forjadev/gun-organization/handler"
	"github.com/forjadev/gun-organization/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Teams = repository.Teams

type TeamService struct {
	repo repository.TeamRepository
}

func NewTeamService(repo repository.TeamRepository) TeamService {
	return TeamService{repo}
}

func (s *TeamService) ListTeams(ctx ...*gin.Context) (Teams, error) {
	if len(ctx) != 1 {
		handler.SendError(ctx[0], http.StatusInternalServerError, "error context length")
		return nil, nil
	}
	teams, err := s.repo.ListTeams()
	if err != nil {
		handler.SendError(ctx[0], http.StatusInternalServerError, "error listing teams")
		return nil, nil
	}
	handler.SendSuccess(ctx[0], "List Teams", teams)
	return teams, nil
}
