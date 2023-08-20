package repository

import (
	"github.com/septian03yogi/model"
	"gorm.io/gorm"
)

type ItemRepository interface {
	BaseRepository[model.Item]
	GetByName(name string) (model.Item, error)
}

type itemRepository struct {
	db *gorm.DB
}

// Create implements ItemRepository.
func (i *itemRepository) Create(payload model.Item) error {
	return i.db.Create(&payload).Error
}

// Delete implements ItemRepository.
func (i *itemRepository) Delete(id string) error {
	var item model.Item
	return i.db.Where("id=?", item.Id).Delete(&item).Error
}

// Get implements ItemRepository.
func (i *itemRepository) Get(id string) (model.Item, error) {
	var item model.Item
	err := i.db.First(&item, "id=?", item.Id).Error
	return item, err
}

// GetByName implements ItemRepository.
func (i *itemRepository) GetByName(name string) (model.Item, error) {
	var item model.Item
	err := i.db.Where("name like $1", "%"+name+"%").Find(&item).Error
	return item, err
}

// List implements ItemRepository.
func (i *itemRepository) List() ([]model.Item, error) {
	var items []model.Item
	err := i.db.Find(&items).Error
	return items, err
}

// Update implements ItemRepository.
func (i *itemRepository) Update(payload model.Item) error {
	return i.db.Model(&payload).Updates(&payload).Error
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}
