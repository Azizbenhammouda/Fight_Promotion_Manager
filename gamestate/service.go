package gamestate

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

const DaysPerTick = 7

type GameStateService interface {
	GetCurrentDate() (time.Time, error)
	Advance() (time.Time, error)
}

type gameStateService struct {
	repo GameStateRepository
}

func NewService(repo GameStateRepository) GameStateService {
	return gameStateService{repo: repo}
}

func (s gameStateService) loadOrCreate() (*GameState, error) {
	state, err := s.repo.Get()
	if err == nil {
		return state, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	state = &GameState{CurrentDate: StartDate}
	if err := s.repo.Save(state); err != nil {
		return nil, err
	}
	return state, nil
}

func (s gameStateService) GetCurrentDate() (time.Time, error) {
	state, err := s.loadOrCreate()
	if err != nil {
		return time.Time{}, err
	}
	return state.CurrentDate, nil
}

func (s gameStateService) Advance() (time.Time, error) {
	state, err := s.loadOrCreate()
	if err != nil {
		return time.Time{}, err
	}

	state.CurrentDate = state.CurrentDate.AddDate(0, 0, DaysPerTick)
	if err := s.repo.Save(state); err != nil {
		return time.Time{}, err
	}
	return state.CurrentDate, nil
}
