package signing

import (
	"github.com/Azizbenhammouda/Fight_Promotion_Manager/fighters"
	"github.com/Azizbenhammouda/Fight_Promotion_Manager/promotions"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	SignFighter(ownerID, fighterID uuid.UUID) error
	MyRoster(ownerID uuid.UUID) ([]fighters.Fighter, error)
}

type signingService struct {
	db         *gorm.DB
	promotions promotions.PromotionService
	fighters   fighters.FighterService
}

func NewService(db *gorm.DB, ps promotions.PromotionService, fs fighters.FighterService) Service {
	return &signingService{db: db, promotions: ps, fighters: fs}
}

func (s *signingService) SignFighter(ownerID, fighterID uuid.UUID) error {
	promo, err := s.promotions.GetMyPromotion(ownerID)
	if err != nil {
		return err
	}
	return Sign(s.db, promo.ID, fighterID)
}

func (s *signingService) MyRoster(ownerID uuid.UUID) ([]fighters.Fighter, error) {
	promo, err := s.promotions.GetMyPromotion(ownerID)
	if err != nil {
		return nil, err
	}
	return s.fighters.ListByPromotion(promo.ID)
}
