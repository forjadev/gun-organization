package repository

import (
	"github.com/forjadev/gun-organization/schemas"
	"gorm.io/gorm"
)

type Teams = []*schemas.Team

// TeamRepository is an interface that defines the methods
// that must be implemented by the TeamRepository
type TeamRepository interface {
	ListTeams() (Teams, error)
}

type TeamRepositoryGorm struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepositoryGorm {
	return TeamRepositoryGorm{db: db}
}

func ListTeams() (Teams, error) {
	var teams Teams
	if err := db.Find(teams).Error; err != nil {
		return nil, err
	}

	return teams, nil
}
