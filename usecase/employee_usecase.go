package usecase

import (
	"fmt"

	"github.com/septian03yogi/model"
	"github.com/septian03yogi/repository"
)

type EmployeeUseCase interface {
	RegisterNewEmployee(payload model.Employee) error
	FindByIdEmployee(id string) (model.Employee, error)
	FindAllEmployee() ([]model.Employee, error)
	UpdateEmployee(payload model.Employee) error
	DeleteEmployee(id string) error
}

type employeeUseCase struct {
	repo repository.EmployeeRepository
}

// DeleteEmployee implements EmployeeUseCase.
func (e *employeeUseCase) DeleteEmployee(id string) error {
	return e.repo.Delete(id)
}

// FindAllEmployee implements EmployeeUseCase.
func (e *employeeUseCase) FindAllEmployee() ([]model.Employee, error) {
	return e.repo.List()
}

// FindByIdEmployee implements EmployeeUseCase.
func (e *employeeUseCase) FindByIdEmployee(id string) (model.Employee, error) {
	return e.repo.Get(id)
}

// RegisterNewEmployee implements EmployeeUseCase.
func (e *employeeUseCase) RegisterNewEmployee(payload model.Employee) error {
	isExistEmployee, _ := e.repo.GetByName(payload.EmployeeName)
	if isExistEmployee.EmployeeName == payload.EmployeeName || payload.EmployeeName == "" {
		return fmt.Errorf("Name required field and can not duplicate %s", payload.EmployeeName)
	}
	err := e.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("Failed to create new employee %v", err)
	}
	return nil
}

// UpdateEmployee implements EmployeeUseCase.
func (e *employeeUseCase) UpdateEmployee(payload model.Employee) error {
	return e.repo.Update(payload)
}

func NewEmployeeUseCase(repo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUseCase{repo: repo}
}
