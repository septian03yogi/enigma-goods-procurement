package repository

import (
	"github.com/septian03yogi/model"
	"gorm.io/gorm"
)

type UomRepository interface {
	BaseRepository[model.Uom]
	GetByName(name string) (model.Uom, error)
}

type uomRepository struct {
	db *gorm.DB
}

// Create implements UomRepository.
func (u *uomRepository) Create(payload model.Uom) error {
	return u.db.Create(&payload).Error
}

// Delete implements UomRepository.
func (u *uomRepository) Delete(id string) error {
	return u.db.Where("id=?", id).Delete(&model.Uom{}).Error
}

// Get implements UomRepository.
func (u *uomRepository) Get(id string) (model.Uom, error) {
	var uom model.Uom
	if err := u.db.Where("id=?", id).First(&uom).Error; err != nil {
		return uom, err
	}
	return uom, nil
}

// GetByName implements UomRepository.
func (u *uomRepository) GetByName(name string) (model.Uom, error) {
	var uom model.Uom
	if err := u.db.Where("uom_name LIKE $1", name).First(&uom).Error; err != nil {
		return uom, err
	}
	return uom, nil
}

// List implements UomRepository.
func (u *uomRepository) List() ([]model.Uom, error) {
	var uoms []model.Uom
	if err := u.db.Find(&uoms).Error; err != nil {
		return uoms, err
	}
	return uoms, nil
}

// Update implements UomRepository.
func (u *uomRepository) Update(payload model.Uom) error {
	return u.db.Model(&payload).Updates(payload).Error
}

func NewUomRepository(db *gorm.DB) UomRepository {
	return &uomRepository{db: db}
}
