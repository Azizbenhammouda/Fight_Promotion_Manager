package promotions

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

const (
	startingBudget     int64 = 50_000
	startingReputation int   = 5
)

type CreatePromotionInput struct {
	Name        string
	HomeCountry Country
}

var ErrInvalidCountry = errors.New("invalid country")
var ErrInvalidName = errors.New("invalid name")
var ErrAlreadyHasPromotion = errors.New("already has a promotion")
var ErrNameTaken = errors.New("name taken")
var ErrPromotionNotFound = errors.New("promotion not found")

type PromotionService interface {
	CreatePromotion(ownerID uuid.UUID, input CreatePromotionInput) (*Promotion, error)
	GetMyPromotion(ownerID uuid.UUID) (*Promotion, error)
}

type promotionService struct {
	repo PromotionRepository
}

func NewPromotionService(repo PromotionRepository) PromotionService {
	return promotionService{
		repo: repo,
	}
}

func (ps promotionService) CreatePromotion(ownerID uuid.UUID, input CreatePromotionInput) (*Promotion, error) {
	name := strings.TrimSpace(input.Name)
	if name == "" {
		return nil, ErrInvalidName
	}
	if !input.HomeCountry.IsValid() {
		return nil, ErrInvalidCountry
	}
	existingPromotion, err := ps.repo.GetByOwnerID(ownerID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingPromotion != nil {
		return nil, ErrAlreadyHasPromotion
	}
	sameName, err := ps.repo.GetByName(name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if sameName != nil {
		return nil, ErrNameTaken
	}
	p := Promotion{
		OwnerID:     ownerID,
		Name:        name,
		HomeCountry: input.HomeCountry,
		Budget:      startingBudget,
		Reputation:  startingReputation,
	}
	err = ps.repo.Create(&p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (ps promotionService) GetMyPromotion(ownerID uuid.UUID) (*Promotion, error) {
	p, err := ps.repo.GetByOwnerID(ownerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPromotionNotFound
		}
		return nil, err
	}
	return p, nil
}
