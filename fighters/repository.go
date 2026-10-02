package fighters

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type FighterRepository interface {
	Create(f *Fighter) error
	GetByID(id uuid.UUID) (*Fighter, error)
	GetByName(name string) (*Fighter, error)
	Update(f *Fighter) error
	Delete(id uuid.UUID) error
}
type fighterRepository struct {
	db *gorm.DB
}

func NewFighterRepository(db *gorm.DB) FighterRepository {
	return fighterRepository{
		db: db,
	}
}
func (fr fighterRepository) Create(f *Fighter) error {
	result := fr.db.Create(f)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (fr fighterRepository) GetByID(id uuid.UUID) (*Fighter, error) {
	var fighter Fighter
	result := fr.db.Where("id = ?", id).First(&fighter)
	if result.Error != nil {
		return nil, result.Error
	}
	return &fighter, nil
}
func (fr fighterRepository) GetByName(name string) (*Fighter, error) {
	var fighter Fighter
	result := fr.db.Where("name = ?", name).First(&fighter)
	if result.Error != nil {
		return nil, result.Error
	}
	return &fighter, nil
}
func (fr fighterRepository) Update(f *Fighter) error {
	result := fr.db.Save(f)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (fr fighterRepository) Delete(id uuid.UUID) error {
	result := fr.db.Where("id = ?", id).Delete(&Fighter{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
