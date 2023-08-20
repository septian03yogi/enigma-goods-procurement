package usecase

import (
	"fmt"

	"github.com/septian03yogi/model"
	"github.com/septian03yogi/repository"
)

type ItemUseCase interface {
	RegisterNewItem(payload model.Item) error
	FindByIdItem(id string) (model.Item, error)
	FindAllItem() ([]model.Item, error)
	DeleteItem(id string) error
	UpdateItem(payload model.Item) error
}

type itemUseCase struct {
	repo repository.ItemRepository
}

// DeleteItem implements ItemUseCase.
func (i *itemUseCase) DeleteItem(id string) error {
	return i.repo.Delete(id)
}

// FindAllItem implements ItemUseCase.
func (i *itemUseCase) FindAllItem() ([]model.Item, error) {
	return i.repo.List()
}

// FindByIdItem implements ItemUseCase.
func (i *itemUseCase) FindByIdItem(id string) (model.Item, error) {
	return i.repo.Get(id)
}

// RegisterNewItem implements ItemUseCase.
func (i *itemUseCase) RegisterNewItem(payload model.Item) error {
	existItem, _ := i.repo.GetByName(payload.ItemName)
	if existItem.ItemName == payload.ItemName || payload.ItemName == "" {
		return fmt.Errorf("Item Name required field and can not duplicate %s", payload.ItemName)
	}
	err := i.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("Failed to create new item %v", err)
	}
	return nil
}

// UpdateItem implements ItemUseCase.
func (i *itemUseCase) UpdateItem(payload model.Item) error {
	return i.repo.Update(payload)
}

func NewItemUseCase(repo repository.ItemRepository) ItemUseCase {
	return &itemUseCase{repo: repo}
}
