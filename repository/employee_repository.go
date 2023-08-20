package repository

import (
	"github.com/septian03yogi/model"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	BaseRepository[model.Employee]
	GetByName(name string) (model.Employee, error)
}

type employeeRepository struct {
	db *gorm.DB
}

// // CekKode implements EmployeeRepository.
// func (e *employeeRepository) CekKode(kode int) (int64, error) {
// 	var count int64
// 	var employee model.Employee
// 	// err := e.db.QueryRow("SELECT COUNT(*) FROM employees WHERE kode=$1", kode).Scan(&count)
// 	err := e.db.Model(&employee).Where("kode=?",kode).Count(&count).Error
// 	if err != nil {
// 		return -1, err
// 	}
// 	return count, nil

// }

// Create implements EmployeeRepository.
func (e *employeeRepository) Create(payload model.Employee) error {
	return e.db.Create(&payload).Error
}

// Delete implements EmployeeRepository.
func (e *employeeRepository) Delete(id string) error {
	var employee model.Employee
	return e.db.Where("id = ?", id).Delete(&employee).Error
}

// Get implements EmployeeRepository.
func (e *employeeRepository) Get(id string) (model.Employee, error) {
	var employee model.Employee
	err := e.db.Where("id= ?", id).First(&employee).Error
	return employee, err
}

// GetByName implements EmployeeRepository.
func (e *employeeRepository) GetByName(name string) (model.Employee, error) {
	var employee model.Employee
	err := e.db.Where("employee_name like $1", "%"+name+"%").Find(&employee).Error
	return employee, err
}

// List implements EmployeeRepository.
func (e *employeeRepository) List() ([]model.Employee, error) {
	var employees []model.Employee
	err := e.db.Find(&employees).Error
	return employees, err
}

// Update implements EmployeeRepository.
func (e *employeeRepository) Update(payload model.Employee) error {
	return e.db.Model(&payload).Updates(payload).Error
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}
