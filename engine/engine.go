package engine

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/Azizbenhammouda/Fight_Promotion_Manager/fighters"
	"github.com/Azizbenhammouda/Fight_Promotion_Manager/gamestate"
)

type Engine struct {
	gameState gamestate.GameStateService
	fighters  fighters.FighterService
	rng       *rand.Rand
}

func New(gs gamestate.GameStateService, fs fighters.FighterService, rng *rand.Rand) *Engine {
	return &Engine{gameState: gs, fighters: fs, rng: rng}
}

type TickResult struct {
	NewDate time.Time `json:"new_date"`
	Events  []string  `json:"events"`
}

func (e *Engine) Tick() (*TickResult, error) {
	newDate, err := e.gameState.Advance()
	if err != nil {
		return nil, err
	}

	active, err := e.fighters.ListActive()
	if err != nil {
		return nil, err
	}

	result := &TickResult{NewDate: newDate, Events: []string{}}

	for _, f := range active {
		age := fighters.AgeOn(f.DateOfBirth, newDate)
		if fighters.ShouldRetire(age, e.rng.Float64()) {
			if err := e.fighters.RetireFighter(f.ID); err != nil {
				return nil, err
			}
			result.Events = append(result.Events,
				fmt.Sprintf("%s (age %d) has retired", f.Name, age))
		}
	}

	return result, nil
}
