package fighters

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type CreateFighterInput struct {
	Name        string
	DateOfBirth time.Time
	Nationality string
}
type FighterService interface {
	CreateFighter(input CreateFighterInput) (*Fighter, error)
	GetFighter(id uuid.UUID) (*Fighter, error)
	DeleteFighter(id uuid.UUID) error
}

var ErrNameTaken = errors.New("name taken")
var ErrFighterNotFound = errors.New("fighter not found")

type fighterService struct {
	repo FighterRepository
}

func NewFighterService(repo FighterRepository) FighterService {
	return fighterService{
		repo: repo,
	}
}

func (fs fighterService) CreateFighter(input CreateFighterInput) (*Fighter, error) {
	existingFighter, err := fs.repo.GetByName(input.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingFighter != nil {
		return nil, ErrNameTaken
	}
	f := Fighter{
		Name:        input.Name,
		DateOfBirth: input.DateOfBirth,
		Nationality: input.Nationality,
	}
	err = fs.repo.Create(&f)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func (fs fighterService) GetFighter(id uuid.UUID) (*Fighter, error) {
	f, err := fs.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFighterNotFound
		}
		return nil, err
	}
	return f, nil
}

func (fs fighterService) DeleteFighter(id uuid.UUID) error {
	err := fs.repo.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrFighterNotFound
		}
		return err
	}
	return nil
}
