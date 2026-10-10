package promotions

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type PromotionRepository interface {
	Create(p *Promotion) error
	GetByID(id uuid.UUID) (*Promotion, error)
	GetByName(name string) (*Promotion, error)
	GetByOwnerID(oid uuid.UUID) (*Promotion, error)
	Update(p *Promotion) error
	Delete(id uuid.UUID) error
}
type promotionRepository struct {
	db *gorm.DB
}

func NewPromotionRepository(db *gorm.DB) PromotionRepository {
	return promotionRepository{
		db: db,
	}
}
func (pr promotionRepository) Create(p *Promotion) error {
	result := pr.db.Create(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (pr promotionRepository) GetByID(id uuid.UUID) (*Promotion, error) {
	var promo Promotion
	result := pr.db.Where("id = ?", id).First(&promo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &promo, nil
}
func (pr promotionRepository) GetByName(name string) (*Promotion, error) {
	var promo Promotion
	result := pr.db.Where("name = ?", name).First(&promo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &promo, nil
}
func (pr promotionRepository) GetByOwnerID(oid uuid.UUID) (*Promotion, error) {
	var promo Promotion
	result := pr.db.Where("owner_id = ?", oid).First(&promo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &promo, nil
}
func (pr promotionRepository) Update(p *Promotion) error {
	result := pr.db.Save(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (pr promotionRepository) Delete(id uuid.UUID) error {
	result := pr.db.Where("id = ?", id).Delete(&Promotion{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
