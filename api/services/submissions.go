package services

import (
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/repository"
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
