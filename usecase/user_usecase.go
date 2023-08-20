package usecase

import (
	"fmt"

	"github.com/septian03yogi/model"
	"github.com/septian03yogi/repository"
)

type UserUseCase interface {
	RegisterNewUser(payload model.UserCredential) error
	FindUserById(id string) (model.UserCredential, error)
	FindAllUser() ([]model.UserCredential, error)
	UpdateUser(payload model.UserCredential) error
	Delete(id string) error
}

type userUseCase struct {
	repo repository.UserRepository
}

// Delete implements UserUseCase.
func (u *userUseCase) Delete(id string) error {
	return u.repo.Delete(id)
}

// FindAllUser implements UserUseCase.
func (u *userUseCase) FindAllUser() ([]model.UserCredential, error) {
	return u.repo.List()
}

// FindUserById implements UserUseCase.
func (u *userUseCase) FindUserById(id string) (model.UserCredential, error) {
	return u.repo.Get(id)
}

// RegisterNewUser implements UserUseCase.
func (u *userUseCase) RegisterNewUser(payload model.UserCredential) error {
	existingUserEmployee, _ := u.repo.GetByIdEmployee(payload.EmployeeId)
	existingUserRole, _ := u.repo.GetByIdRole(payload.RoleId)
	if existingUserEmployee.EmployeeId == payload.EmployeeId && existingUserRole.RoleId == payload.RoleId {
		return fmt.Errorf("User with Employee ID %s and Role ID %s already exist", payload.EmployeeId, payload.RoleId)
	}
	err := u.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("Failed to create New User %v", err)
	}
	return nil
}

// UpdateUser implements UserUseCase.
func (u *userUseCase) UpdateUser(payload model.UserCredential) error {
	return u.repo.Update(payload)
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
