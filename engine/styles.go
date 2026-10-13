package engine

import (
	"github.com/Azizbenhammouda/Fight_Promotion_Manager/fighters"
	"github.com/Azizbenhammouda/Fight_Promotion_Manager/promotions"
)

var nationalityStyles = map[promotions.Country]fighters.FightingSkills{
	promotions.Thailand: {MuayThai: 20},
	promotions.Russia:   {Wrestling: 20},
	promotions.USA:      {Boxing: 20},
	promotions.Brazil:   {BJJ: 20},
}
