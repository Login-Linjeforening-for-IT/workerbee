package services

import (
	"workerbee/models"
	"workerbee/repositories"
)

type SubmissionService struct {
	repo repositories.Submissionrepositories
}

func NewSubmissionService(repo repositories.Submissionrepositories) *SubmissionService {
	return &SubmissionService{repo: repo}
}

func (s *SubmissionService) GetSubmission(formID, submissionID string) (*models.Submission, error) {
	return s.repo.GetSubmission(formID, submissionID)
}
