package repository

import (
	"github.com/septian03yogi/model"
	"gorm.io/gorm"
)

type RoleUserRepository interface {
	BaseRepository[model.RoleUser]
	GetByName(name string) (model.RoleUser, error)
}

type roleUserRepository struct {
	db *gorm.DB
}

// Create implements RoleUserRepository.
func (r *roleUserRepository) Create(payload model.RoleUser) error {
	err := r.db.Create(&payload).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete implements RoleUserRepository.
func (r *roleUserRepository) Delete(id string) error {
	var role model.RoleUser
	err := r.db.Where("id=?", id).Delete(&role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get implements RoleUserRepository.
func (r *roleUserRepository) Get(id string) (model.RoleUser, error) {
	var role model.RoleUser
	err := r.db.Where("id=?", id).First(&role).Error
	if err != nil {
		return role, err
	}
	return role, nil
}

// GetByName implements RoleUserRepository.
func (r *roleUserRepository) GetByName(name string) (model.RoleUser, error) {
	var role model.RoleUser
	if err := r.db.Where("role_name like $1", name).First(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}

// List implements RoleUserRepository.
func (r *roleUserRepository) List() ([]model.RoleUser, error) {
	var roles []model.RoleUser
	if err := r.db.Find(&roles).Error; err != nil {
		return roles, err
	}
	return roles, nil
}

// Update implements RoleUserRepository.
func (r *roleUserRepository) Update(payload model.RoleUser) error {
	return r.db.Model(&payload).Updates(&payload).Error
}

func NewRoleUserRepository(db *gorm.DB) RoleUserRepository {
	return &roleUserRepository{db: db}
}
