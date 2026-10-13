package fighters

import (
	"errors"
	"strings"
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
	ListActive() ([]Fighter, error)
	RetireFighter(id uuid.UUID) error
	CreateGeneratedFighter(f *Fighter) (*Fighter, error)
}

var ErrInvalidInput = errors.New("invalid input")
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
	input.Name = strings.TrimSpace(input.Name)
	input.Nationality = strings.ToLower(strings.TrimSpace(input.Nationality))
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
func (fs fighterService) ListActive() ([]Fighter, error) {
	return fs.repo.ListActive()
}

func (fs fighterService) RetireFighter(id uuid.UUID) error {
	f, err := fs.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrFighterNotFound
		}
		return err
	}
	f.Retired = true
	return fs.repo.Update(f)
}

func (fs fighterService) CreateGeneratedFighter(f *Fighter) (*Fighter, error) {
	f.Name = strings.TrimSpace(f.Name)
	f.Nationality = strings.ToLower(strings.TrimSpace(f.Nationality))
	if f.Name == "" || f.Nationality == "" {
		return nil, ErrInvalidInput
	}

	existing, err := fs.repo.GetByName(f.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existing != nil {
		return nil, ErrNameTaken
	}

	if err := fs.repo.Create(f); err != nil {
		return nil, err
	}
	return f, nil
}
