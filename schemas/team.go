package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name string
}

type TeamNameRequest struct {
	name string
}

type TeamResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	Name      string    `json:"name"`
}
