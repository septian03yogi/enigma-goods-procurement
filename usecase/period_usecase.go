package usecase

import (
	"fmt"

	"github.com/septian03yogi/model"
	"github.com/septian03yogi/repository"
)

type PeriodUseCase interface {
	RegisterNewPeriod(payload model.Period) error
	FindByIdPeriod(id string) (model.Period, error)
	FindAllPeriod() ([]model.Period, error)
	UpdatePeriod(payload model.Period) error
	DeletePeriod(id string) error
}

type periodUseCase struct {
	repo repository.PeriodRepository
}

// DeletePeriod implements PeriodUseCase.
func (p *periodUseCase) DeletePeriod(id string) error {
	return p.repo.Delete(id)
}

// FindAllPeriod implements PeriodUseCase.
func (p *periodUseCase) FindAllPeriod() ([]model.Period, error) {
	return p.repo.List()
}

// FindByIdPeriod implements PeriodUseCase.
func (p *periodUseCase) FindByIdPeriod(id string) (model.Period, error) {
	return p.repo.Get(id)
}

// RegisterNewPeriod implements PeriodUseCase.
func (p *periodUseCase) RegisterNewPeriod(payload model.Period) error {
	existingPeriod, _ := p.repo.GetByName(payload.PeriodName)
	if existingPeriod.PeriodName == payload.PeriodName || payload.PeriodName == "" {
		return fmt.Errorf("Period Name require field and can not duplicate %s", payload.PeriodName)
	}
	err := p.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("Failed to create new period %v", err)
	}
	return nil
}

// UpdatePeriod implements PeriodUseCase.
func (p *periodUseCase) UpdatePeriod(payload model.Period) error {
	return p.repo.Update(payload)
}

func NewPeriodUseCase(repo repository.PeriodRepository) PeriodUseCase {
	return &periodUseCase{repo: repo}
}
