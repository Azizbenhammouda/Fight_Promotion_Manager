package fighters

import (
	"time"

	"github.com/google/uuid"
)

type FightingSkills struct {
	Boxing     int `gorm:"column:boxing"`
	KickBoxing int `gorm:"column:kick_boxing"`
	MuayThai   int `gorm:"column:muay_thai"`
	Wrestling  int `gorm:"column:wrestling"`
	BJJ        int `gorm:"column:bjj"`
	Judo       int `gorm:"column:judo"`
}
type Fighter struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"unique;not null"`
	DateOfBirth time.Time      `json:"date_of_birth"`
	Promotion   string         `json:"promotion"` // to be chnaged later by maybe promotion.id
	Nationality string         `json:"nationality"`
	Skills      FightingSkills `gorm:"embedded"`
	Personality string         `json:"personality"`
	Popularity  int            `json:"popularity"`
}
