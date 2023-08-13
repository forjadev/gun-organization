package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	GithubID string
	Login    string
	URL      string
	Role     string
	TeamID   int
	Team     Team
}

type MemberResponse struct {
	ID        uint      `json:"id"`
	GithubID  string    `json:"github_id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	Login     string    `json:"login"`
	URL       string    `json:"html_url"`
	Role      string    `json:"role"`
	Team      Team      `json:"team"`
}
