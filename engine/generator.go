package engine

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/Azizbenhammouda/Fight_Promotion_Manager/fighters"
	"github.com/Azizbenhammouda/Fight_Promotion_Manager/promotions"
	"github.com/google/uuid"
)

const (
	minGeneratedAge   = 18
	maxGeneratedAge   = 30
	baseSkillMin      = 10
	baseSkillMax      = 40
	maxNameAttempts   = 10
	starterSkillMin   = 5
	starterSkillMax   = 25
	starterRosterSize = 4
)

var countries = []promotions.Country{
	promotions.Thailand, promotions.Russia, promotions.USA, promotions.Brazil,
}

var firstNames = map[promotions.Country][]string{
	promotions.Thailand: {"Somchai", "Anan", "Kritsada", "Niran", "Prasert", "Chai", "Surin", "Taksin"},
	promotions.Russia:   {"Ivan", "Dmitri", "Sergei", "Alexei", "Nikolai", "Pavel", "Viktor", "Oleg"},
	promotions.USA:      {"Jake", "Tyler", "Marcus", "Cody", "Brandon", "Dwayne", "Colt", "Travis"},
	promotions.Brazil:   {"Thiago", "Rafael", "Lucas", "Mateus", "Diego", "Gabriel", "Caio", "Bruno"},
}

var lastNames = map[promotions.Country][]string{
	promotions.Thailand: {"Srisuk", "Kaewmanee", "Boonmee", "Chaiyasit", "Wongsawat", "Phanit", "Rattana", "Saetang"},
	promotions.Russia:   {"Petrov", "Volkov", "Sokolov", "Morozov", "Kuznetsov", "Orlov", "Fedorov", "Gromov"},
	promotions.USA:      {"Carter", "Brooks", "Hayes", "Mercer", "Dawson", "Reed", "Walker", "Stone"},
	promotions.Brazil:   {"Silva", "Santos", "Oliveira", "Costa", "Pereira", "Almeida", "Ribeiro", "Moura"},
}

var personalities = []string{"aggressive", "calm", "arrogant", "humble", "reckless"}

func pick[T any](rng *rand.Rand, list []T) T {
	return list[rng.IntN(len(list))]
}

func (e *Engine) baseSkill(lo, hi int) int {
	return lo + e.rng.IntN(hi-lo+1)
}

func addSkills(a, b fighters.FightingSkills) fighters.FightingSkills {
	return fighters.FightingSkills{
		Boxing:     a.Boxing + b.Boxing,
		KickBoxing: a.KickBoxing + b.KickBoxing,
		MuayThai:   a.MuayThai + b.MuayThai,
		Wrestling:  a.Wrestling + b.Wrestling,
		BJJ:        a.BJJ + b.BJJ,
		Judo:       a.Judo + b.Judo,
	}
}

func (e *Engine) GenerateFreeAgent(now time.Time) (*fighters.Fighter, error) {
	return e.generateFighter(now, pick(e.rng, countries), nil, baseSkillMin, baseSkillMax)
}

func (e *Engine) generateFighter(now time.Time, country promotions.Country, promotionID *uuid.UUID, lo, hi int) (*fighters.Fighter, error) {
	age := minGeneratedAge + e.rng.IntN(maxGeneratedAge-minGeneratedAge+1)
	dob := now.AddDate(-age, 0, -e.rng.IntN(365))

	base := fighters.FightingSkills{
		Boxing:     e.baseSkill(lo, hi),
		KickBoxing: e.baseSkill(lo, hi),
		MuayThai:   e.baseSkill(lo, hi),
		Wrestling:  e.baseSkill(lo, hi),
		BJJ:        e.baseSkill(lo, hi),
		Judo:       e.baseSkill(lo, hi),
	}
	skills := addSkills(base, nationalityStyles[country])

	for attempt := 0; attempt < maxNameAttempts; attempt++ {
		f := &fighters.Fighter{
			Name:        pick(e.rng, firstNames[country]) + " " + pick(e.rng, lastNames[country]),
			DateOfBirth: dob,
			Nationality: string(country),
			PromotionID: promotionID,
			Skills:      skills,
			Personality: pick(e.rng, personalities),
		}
		created, err := e.fighters.CreateGeneratedFighter(f)
		if errors.Is(err, fighters.ErrNameTaken) {
			continue
		}
		if err != nil {
			return nil, err
		}
		return created, nil
	}
	return nil, fmt.Errorf("no free name after %d attempts", maxNameAttempts)
}
func (e *Engine) GenerateStarterRoster(promotionID uuid.UUID, country promotions.Country) error {
	now, err := e.gameState.GetCurrentDate()
	if err != nil {
		return err
	}
	for i := 0; i < starterRosterSize; i++ {
		if _, err := e.generateFighter(now, country, &promotionID, starterSkillMin, starterSkillMax); err != nil {
			return err
		}
	}
	return nil
}
