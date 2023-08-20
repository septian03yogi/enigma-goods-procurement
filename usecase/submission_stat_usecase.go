package usecase

import (
	"fmt"

	"github.com/septian03yogi/model"
	"github.com/septian03yogi/repository"
)

type SubmisisonStatusUseCase interface {
	RegisterNewSubStatus(payload model.SubmisisonStatus) error
	FindSubStatusById(id string) (model.SubmisisonStatus, error)
	FindAllSubStatus() ([]model.SubmisisonStatus, error)
	UpdateSubStatus(payload model.SubmisisonStatus) error
	DeleteSubStatus(id string) error
}

type submissionStatusUseCase struct {
	repo repository.SubmissionStatusRepository
}

// DeleteSubStatus implements SubmisisonStatusUseCase.
func (s *submissionStatusUseCase) DeleteSubStatus(id string) error {
	return s.repo.Delete(id)
}

// FindAllSubStatus implements SubmisisonStatusUseCase.
func (s *submissionStatusUseCase) FindAllSubStatus() ([]model.SubmisisonStatus, error) {
	return s.repo.List()
}

// FindSubStatusById implements SubmisisonStatusUseCase.
func (s *submissionStatusUseCase) FindSubStatusById(id string) (model.SubmisisonStatus, error) {
	return s.repo.Get(id)
}

// RegisterNewSubStatus implements SubmisisonStatusUseCase.
func (s *submissionStatusUseCase) RegisterNewSubStatus(payload model.SubmisisonStatus) error {
	existingStatus, _ := s.repo.GetByDetail(payload.StatusDetail)
	if existingStatus.StatusDetail == payload.StatusDetail || payload.StatusDetail == "" {
		return fmt.Errorf("Status Detail require field and can not duplicate %s", payload.StatusDetail)
	}
	err := s.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("Failed to create new submission status %v", err)
	}
	return nil
}

// UpdateSubStatus implements SubmisisonStatusUseCase.
func (s *submissionStatusUseCase) UpdateSubStatus(payload model.SubmisisonStatus) error {
	return s.repo.Update(payload)
}

func NewSubmissionStatusUseCase(repo repository.SubmissionStatusRepository) SubmisisonStatusUseCase {
	return &submissionStatusUseCase{repo: repo}
}
