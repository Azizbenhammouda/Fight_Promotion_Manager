package gamestate

import "gorm.io/gorm"

type GameStateRepository interface {
	Get() (*GameState, error)
	Save(state *GameState) error
}

type gameStateRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) GameStateRepository {
	return gameStateRepository{db: db}
}

func (r gameStateRepository) Get() (*GameState, error) {
	var state GameState
	if err := r.db.First(&state).Error; err != nil {
		return nil, err
	}
	return &state, nil
}

func (r gameStateRepository) Save(state *GameState) error {
	return r.db.Save(state).Error
}
