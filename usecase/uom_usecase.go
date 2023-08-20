package usecase

import (
	"fmt"

	"github.com/septian03yogi/model"
	"github.com/septian03yogi/repository"
)

type UomUseCase interface {
	RegisterNewUom(payload model.Uom) error
	FindUomById(id string) (model.Uom, error)
	FindAllUom() ([]model.Uom, error)
	UpdateUom(payload model.Uom) error
	DeleteUom(id string) error
}

type uomUseCase struct {
	repo repository.UomRepository
}

// DeleteUom implements UomUseCase.
func (u *uomUseCase) DeleteUom(id string) error {
	return u.repo.Delete(id)
}

// FindAllUom implements UomUseCase.
func (u *uomUseCase) FindAllUom() ([]model.Uom, error) {
	return u.repo.List()
}

// FindUomById implements UomUseCase.
func (u *uomUseCase) FindUomById(id string) (model.Uom, error) {
	return u.repo.Get(id)
}

// RegisterNewUom implements UomUseCase.
func (u *uomUseCase) RegisterNewUom(payload model.Uom) error {
	existingUom, _ := u.repo.GetByName(payload.UomName)
	if existingUom.UomName == payload.UomName || payload.UomName == "" {
		return fmt.Errorf("Name require field and should not duplicate %s", payload.UomName)
	}
	err := u.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("Failed to create new uom %v", err)
	}
	return nil
}

// UpdateUom implements UomUseCase.
func (u *uomUseCase) UpdateUom(payload model.Uom) error {
	return u.repo.Update(payload)
}

func NewUomUseCase(repo repository.UomRepository) UomUseCase {
	return &uomUseCase{repo: repo}
}
