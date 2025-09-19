package services

import (
	"workerbee/models"
	"workerbee/repository"
)

type SubmissionService struct {
	repo repository.SubmissionRepository
}

func NewSubmissionService(repo repository.SubmissionRepository) *SubmissionService {
	return &SubmissionService{repo: repo}
}

func (s *SubmissionService) GetSubmission(formID, submissionID string) (*models.Submission, error) {
	return s.repo.GetSubmission(formID, submissionID)
}
