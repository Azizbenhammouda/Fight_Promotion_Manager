package promotions

import (
	"time"

	"github.com/google/uuid"
)

type Country string

const (
	Thailand Country = "thailand"
	Russia   Country = "russia"
	USA      Country = "usa"
	Brazil   Country = "brazil"
)

func (c Country) IsValid() bool {
	switch c {
	case Thailand, Russia, USA, Brazil:
		return true
	}
	return false
}

type Promotion struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	OwnerID     uuid.UUID `json:"owner_id" gorm:"type:uuid;not null;uniqueIndex"`
	Name        string    `json:"name" gorm:"unique;not null"`
	HomeCountry Country   `json:"home_country"`
	Budget      int64     `json:"budget"`
	Reputation  int       `json:"reputation"`
	CreatedAt   time.Time `json:"created_at"`
}
