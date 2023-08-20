package repository

import (
	"github.com/septian03yogi/model"
	"gorm.io/gorm"
)

type PeriodRepository interface {
	BaseRepository[model.Period]
	GetByName(name string) (model.Period, error)
}

type periodRepository struct {
	db *gorm.DB
}

// Create implements PeriodRepository.
func (p *periodRepository) Create(payload model.Period) error {
	return p.db.Create(&payload).Error
}

// Delete implements PeriodRepository.
func (p *periodRepository) Delete(id string) error {
	var period model.Period
	err := p.db.Where("id=?", id).Delete(&period).Error
	return err
}

// Get implements PeriodRepository.
func (p *periodRepository) Get(id string) (model.Period, error) {
	var period model.Period
	err := p.db.Where("id=?", id).First(&period).Error
	return period, err
}

// GetByName implements PeriodRepository.
func (p *periodRepository) GetByName(name string) (model.Period, error) {
	var period model.Period
	err := p.db.Where("period_name like $1", name).First(&period).Error
	return period, err
}

// List implements PeriodRepository.
func (p *periodRepository) List() ([]model.Period, error) {
	var periods []model.Period
	err := p.db.Find(&periods).Error
	return periods, err
}

// Update implements PeriodRepository.
func (p *periodRepository) Update(payload model.Period) error {
	return p.db.Model(&payload).Updates(&payload).Error
}

func NewPeriodRepository(db *gorm.DB) PeriodRepository {
	return &periodRepository{db: db}
}
