package repository

import (
	"github.com/septian03yogi/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[model.UserCredential]
	GetByIdEmployee(id string) (model.UserCredential, error)
	GetByIdRole(id string) (model.UserCredential, error)
}

type userRepository struct {
	db *gorm.DB
}

// GetByIdRole implements UserRepository.
func (u *userRepository) GetByIdRole(id string) (model.UserCredential, error) {
	var user model.UserCredential
	if err := u.db.Where("role_id=?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Create implements UserRepository.
func (u *userRepository) Create(payload model.UserCredential) error {
	return u.db.Create(&payload).Error
}

// Delete implements UserRepository.
func (u *userRepository) Delete(id string) error {
	var user model.UserCredential
	return u.db.Where("id=?", id).Delete(&user).Error
}

// Get implements UserRepository.
func (u *userRepository) Get(id string) (model.UserCredential, error) {
	var user model.UserCredential
	if err := u.db.Where("id=?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetByIdEmployee implements UserRepository.
func (u *userRepository) GetByIdEmployee(id string) (model.UserCredential, error) {
	var user model.UserCredential
	if err := u.db.Where("employee_id=?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// List implements UserRepository.
func (u *userRepository) List() ([]model.UserCredential, error) {
	var users []model.UserCredential
	if err := u.db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

// Update implements UserRepository.
func (u *userRepository) Update(payload model.UserCredential) error {
	return u.db.Model(&payload).Updates(payload).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
