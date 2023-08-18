package usecase

import (
	"fmt"

	"github.com/septian03yogi/model"
	"github.com/septian03yogi/repository"
)

type DepartmentUseCase interface {
	RegisterNewDepartment(payload model.Department) error
	FindByIdDepartment(id string) (model.Department, error)
	FindAllDepartment() ([]model.Department, error)
	UpdateDepartment(payload model.Department) error
	DeleteDepartment(id string) error
}

type departmentUseCase struct {
	repo repository.DepartmentRepository
}

// RegisterNewDepartment implements DepartmentUseCase.
func (d *departmentUseCase) RegisterNewDepartment(payload model.Department) error {
	isExistDepartmentName, _ := d.repo.GetByName(payload.DepartmentName)
	if payload.DepartmentName == "" || isExistDepartmentName.DepartmentName == payload.DepartmentName {
		return fmt.Errorf("Department name required field and can not duplicate %s", payload.DepartmentName)
	}
	err := d.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("Failed to create new Department %v", err)
	}
	return nil
}

// DeleteDepartment implements DepartmentUseCase.
func (d *departmentUseCase) DeleteDepartment(id string) error {
	return d.repo.Delete(id)
}

// FindAllDepartment implements DepartmentUseCase.
func (d *departmentUseCase) FindAllDepartment() ([]model.Department, error) {
	return d.repo.List()
}

// FindByIdDepartment implements DepartmentUseCase.
func (d *departmentUseCase) FindByIdDepartment(id string) (model.Department, error) {
	return d.repo.Get(id)
}

// UpdateDepartment implements DepartmentUseCase.
func (d *departmentUseCase) UpdateDepartment(payload model.Department) error {
	return d.repo.Update(payload)
}

func NewDepartmentUseCase(repo repository.DepartmentRepository) DepartmentUseCase {
	return &departmentUseCase{repo: repo}
}
