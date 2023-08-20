package repository

import (
	"github.com/septian03yogi/model"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	BaseRepository[model.Department]
	GetByName(DepartmentName string) (model.Department, error)
}

type departmentRepository struct {
	db *gorm.DB
}

// GetByName implements DepartmentRepository.
func (d *departmentRepository) GetByName(departmentName string) (model.Department, error) {
	var department model.Department
	err := d.db.Where("department_name LIKE $1", "%"+departmentName+"%").Find(&department).Error
	return department, err
}

// Create implements DepartmentRepository.
func (d *departmentRepository) Create(payload model.Department) error {
	return d.db.Create(&payload).Error
}

// Delete implements DepartmentRepository.
func (d *departmentRepository) Delete(id string) error {
	department := model.Department{}
	return d.db.Where("id = ?", id).Delete(&department).Error
}

// Get implements DepartmentRepository.
func (d *departmentRepository) Get(id string) (model.Department, error) {
	var department model.Department
	err := d.db.Where("id= ?", id).First(&department).Error
	return department, err
}

// List implements DepartmentRepository.
func (d *departmentRepository) List() ([]model.Department, error) {
	var departments []model.Department
	result := d.db.Find(&departments)
	if result.Error != nil {
		return nil, result.Error
	}
	return departments, result.Error
}

// Update implements DepartmentRepository.
func (d *departmentRepository) Update(payload model.Department) error {
	return d.db.Model(&payload).Updates(payload).Error
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db: db}
}
