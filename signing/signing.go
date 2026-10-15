package signing

import (
	"errors"

	"github.com/Azizbenhammouda/Fight_Promotion_Manager/fighters"
	"github.com/Azizbenhammouda/Fight_Promotion_Manager/promotions"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrFighterRetired     = errors.New("fighter retired")
	ErrFighterNotFree     = errors.New("fighter already has a promotion")
	ErrInsufficientBudget = errors.New("not enough budget")
)

func Sign(db *gorm.DB, promotionID, fighterID uuid.UUID) error {
	return db.Transaction(func(tx *gorm.DB) error {
		fr := fighters.NewFighterRepository(tx)
		pr := promotions.NewPromotionRepository(tx)

		fighter, err := fr.GetByID(fighterID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fighters.ErrFighterNotFound
			}
			return err
		}
		if fighter.Retired {
			return ErrFighterRetired
		}
		if fighter.PromotionID != nil {
			return ErrFighterNotFree
		}

		promo, err := pr.GetByID(promotionID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return promotions.ErrPromotionNotFound
			}
			return err
		}
		fee := MonthlySalary(fighter.Skills, fighter.Popularity, promo.Reputation)
		if promo.Budget < fee {
			return ErrInsufficientBudget
		}

		if err := tx.Model(promo).Update("budget", promo.Budget-fee).Error; err != nil {
			return err
		}
		return tx.Model(fighter).Update("promotion_id", promo.ID).Error
	})
}
