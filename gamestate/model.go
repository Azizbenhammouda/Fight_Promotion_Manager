package gamestate

import (
	"time"

	"github.com/google/uuid"
)

var StartDate = time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)

type GameState struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CurrentDate time.Time `json:"current_date" gorm:"not null"`
}
