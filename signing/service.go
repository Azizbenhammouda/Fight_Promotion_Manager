package signing

import (
	"github.com/Azizbenhammouda/Fight_Promotion_Manager/promotions"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service interface {
	SignFighter(ownerID, fighterID uuid.UUID) error
}

type signingService struct {
	db         *gorm.DB
	promotions promotions.PromotionService
}

func NewService(db *gorm.DB, ps promotions.PromotionService) Service {
	return &signingService{db: db, promotions: ps}
}

func (s *signingService) SignFighter(ownerID, fighterID uuid.UUID) error {
	promo, err := s.promotions.GetMyPromotion(ownerID)
	if err != nil {
		return err
	}
	return Sign(s.db, promo.ID, fighterID)
}
