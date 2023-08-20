package repository

import (
	"github.com/septian03yogi/model"
	"gorm.io/gorm"
)

type SubmissionRepository interface {
	BaseRepository[model.Submission]
	GetByIdEmployee(idEmployee string) (model.Submission, error)
	GetByIdPeriod(idPeriod string) (model.Submission, error)
}

type submissionRepository struct {
	db *gorm.DB
}

// Create implements SubmissionRepository.
func (s *submissionRepository) Create(payload model.Submission) error {
	return s.db.Create(&payload).Error
}

// Delete implements SubmissionRepository.
func (s *submissionRepository) Delete(id string) error {
	var submission model.Submission
	return s.db.Where("id=?", id).Delete(&submission).Error
}

// Get implements SubmissionRepository.
func (s *submissionRepository) Get(id string) (model.Submission, error) {
	var submission model.Submission
	err := s.db.Where("id=?", id).First(&submission).Error
	if err != nil {
		return submission, err
	}
	return submission, nil
}

// GetByIdEmployee implements SubmissionRepository.
func (s *submissionRepository) GetByIdEmployee(idEmployee string) (model.Submission, error) {
	var submission model.Submission
	if err := s.db.Where("employee_id=?", idEmployee).First(&submission).Error; err != nil {
		return submission, err
	}
	return submission, nil
}

// GetByIdPeriod implements SubmissionRepository.
func (s *submissionRepository) GetByIdPeriod(idPeriod string) (model.Submission, error) {
	var submission model.Submission
	if err := s.db.Where("period_id=?", idPeriod).First(&submission).Error; err != nil {
		return submission, err
	}
	return submission, nil
}

// List implements SubmissionRepository.
func (s *submissionRepository) List() ([]model.Submission, error) {
	var submissions []model.Submission
	if err := s.db.Find(&submissions).Error; err != nil {
		return submissions, err
	}
	return submissions, nil
}

// Update implements SubmissionRepository.
func (s *submissionRepository) Update(payload model.Submission) error {
	if err := s.db.Model(&payload).Updates(payload).Error; err != nil {
		return err
	}
	return nil
}

func NewSubmissionRepository(db *gorm.DB) SubmissionRepository {
	return &submissionRepository{db: db}
}
