package usecase

import (
	"fmt"

	"github.com/septian03yogi/model"
	"github.com/septian03yogi/repository"
)

type SubmissionUseCase interface {
	RegisterNewSubmission(payload model.Submission) error
	FindSubmissionById(id string) (model.Submission, error)
	FindAllSubmission() ([]model.Submission, error)
	UpdateSubmission(payload model.Submission) error
	DeleteSubmission(id string) error
}

type submissionUseCase struct {
	repo repository.SubmissionRepository
}

// DeleteSubmission implements SubmissionUseCase.
func (s *submissionUseCase) DeleteSubmission(id string) error {
	return s.repo.Delete(id)
}

// FindAllSubmission implements SubmissionUseCase.
func (s *submissionUseCase) FindAllSubmission() ([]model.Submission, error) {
	return s.repo.List()
}

// FindSubmissionById implements SubmissionUseCase.
func (s *submissionUseCase) FindSubmissionById(id string) (model.Submission, error) {
	return s.repo.Get(id)
}

// RegisterNewSubmission implements SubmissionUseCase.
func (s *submissionUseCase) RegisterNewSubmission(payload model.Submission) error {
	// existingEmployee, _ := s.repo.GetByIdEmployee(payload.EmployeeId)
	// existingPeriod, _ := s.repo.GetByIdPeriod(payload.PeriodId)
	// if existingEmployee.EmployeeId == payload.EmployeeId && existingPeriod.PeriodId == payload.PeriodId {
	// 	return fmt.Errorf("Employee with ID %s can not submit in the same period ID %s", payload.EmployeeId, payload.PeriodId)
	// }

	err := s.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("Failed to create new submission %v", err)
	}
	return nil
}

// UpdateSubmission implements SubmissionUseCase.
func (s *submissionUseCase) UpdateSubmission(payload model.Submission) error {
	return s.repo.Update(payload)
}

func NewSubmissionUseCase(repo repository.SubmissionRepository) SubmissionUseCase {
	return &submissionUseCase{repo: repo}
}
