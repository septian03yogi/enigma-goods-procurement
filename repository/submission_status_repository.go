package repository

import (
	"github.com/septian03yogi/model"
	"gorm.io/gorm"
)

type SubmissionStatusRepository interface {
	BaseRepository[model.SubmisisonStatus]
	GetByDetail(detail string) (model.SubmisisonStatus, error)
}

type submissionStatusRepository struct {
	db *gorm.DB
}

// GetByDetail implements SubmissionStatusRepository.
func (s *submissionStatusRepository) GetByDetail(detail string) (model.SubmisisonStatus, error) {
	var status model.SubmisisonStatus
	if err := s.db.Where("status_detail LIKE $1", detail).First(&status).Error; err != nil {
		return status, err
	}
	return status, nil
}

// Create implements SubmissionStatusRepository.
func (s *submissionStatusRepository) Create(payload model.SubmisisonStatus) error {
	return s.db.Create(&payload).Error
}

// Delete implements SubmissionStatusRepository.
func (s *submissionStatusRepository) Delete(id string) error {
	var submissionStatus model.SubmisisonStatus
	return s.db.Where("id=?", id).Delete(&submissionStatus).Error
}

// Get implements SubmissionStatusRepository.
func (s *submissionStatusRepository) Get(id string) (model.SubmisisonStatus, error) {
	var submissionStatus model.SubmisisonStatus
	if err := s.db.Where("id=?", id).First(&submissionStatus).Error; err != nil {
		return submissionStatus, err
	}
	return submissionStatus, nil
}

// List implements SubmissionStatusRepository.
func (s *submissionStatusRepository) List() ([]model.SubmisisonStatus, error) {
	var submissionStatus []model.SubmisisonStatus
	if err := s.db.Find(&submissionStatus).Error; err != nil {
		return submissionStatus, err
	}
	return submissionStatus, nil
}

// Update implements SubmissionStatusRepository.
func (s *submissionStatusRepository) Update(payload model.SubmisisonStatus) error {
	return s.db.Model(&payload).Updates(payload).Error
}

func NewSubmissionStatusRepository(db *gorm.DB) SubmissionStatusRepository {
	return &submissionStatusRepository{db: db}
}
