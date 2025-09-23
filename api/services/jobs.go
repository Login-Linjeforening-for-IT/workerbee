package services

import (
	"workerbee/models"
	"workerbee/repository"
)

type JobsService struct {
	repo repository.JobsRepository
}

func NewJobsService(repo repository.JobsRepository) *JobsService {
	return &JobsService{repo: repo}
}

func (s *JobsService) GetJobs(search, limit, offset string) ([]models.JobWithTotalCount, error) {
	return s.repo.GetJobs(search, limit, offset)
}