package usecase

import (
	"fmt"

	"github.com/septian03yogi/model"
	"github.com/septian03yogi/repository"
)

type RoleUserUseCase interface {
	RegisterNewRole(payload model.RoleUser) error
	FindByIdRole(id string) (model.RoleUser, error)
	FindAllRole() ([]model.RoleUser, error)
	UpdateRole(payload model.RoleUser) error
	DeleteRole(id string) error
}

type roleUserUseCase struct {
	repo repository.RoleUserRepository
}

// DeleteRole implements RoleUserUseCase.
func (r *roleUserUseCase) DeleteRole(id string) error {
	return r.repo.Delete(id)
}

// FindAllRole implements RoleUserUseCase.
func (r *roleUserUseCase) FindAllRole() ([]model.RoleUser, error) {
	return r.repo.List()
}

// FindByIdRole implements RoleUserUseCase.
func (r *roleUserUseCase) FindByIdRole(id string) (model.RoleUser, error) {
	return r.repo.Get(id)
}

// RegisterNewRole implements RoleUserUseCase.
func (r *roleUserUseCase) RegisterNewRole(payload model.RoleUser) error {
	existingRole, _ := r.repo.GetByName(payload.RoleName)
	if existingRole.RoleName == payload.RoleName || payload.RoleName == "" {
		return fmt.Errorf("Role Name require field and can not duplicate %s", payload.RoleName)
	}
	err := r.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("Failed to create new role user %v", err)
	}
	return nil
}

// UpdateRole implements RoleUserUseCase.
func (r *roleUserUseCase) UpdateRole(payload model.RoleUser) error {
	return r.repo.Update(payload)
}

func NewRoleUserUseCase(repo repository.RoleUserRepository) RoleUserUseCase {
	return &roleUserUseCase{repo: repo}
}
