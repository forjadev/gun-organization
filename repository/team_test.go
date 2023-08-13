package repository

import (
	"github.com/forjadev/gun-organization/service"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamsRoute(t *testing.T) {
	t.Run("load teams path", func(t *testing.T) {
		router := _setupTestRouter()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/teams", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestTeamService_ListTeam(t *testing.T) {
	t.Run("receive all users", func(t *testing.T) {
		// Setup
		mockRepo := &MockTeamRepository{
			ListTeamsFn: func() (Teams, error) {
				return Teams{
					{Model: gorm.Model{ID: 1}, Name: "Mock User"},
					{Model: gorm.Model{ID: 2}, Name: "Mock Mock"},
				}, nil
			},
		}
		teamService := service.NewTeamService(mockRepo)

		// Execute
		teams, err := teamService.ListTeams()

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, teams)
		assert.Equal(t, "Mock User", teams[1].Name)
		assert.Equal(t, "Mock Mock", teams[2].Name)
	})
}

// MockTeamRepository is a mock of TeamRepository interface
// that will be used on TeamService ListTeams method
type MockTeamRepository struct {
	ListTeamsFn func() (Teams, error)
}

func (m *MockTeamRepository) ListTeams() (Teams, error) {
	return m.ListTeamsFn()
}
