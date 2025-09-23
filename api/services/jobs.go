package services

import (
	"strconv"
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repository"
)

type JobsService struct {
	repo repository.JobsRepository
}

func NewJobsService(repo repository.JobsRepository) *JobsService {
	return &JobsService{repo: repo}
}

func (s *JobsService) GetJobs(search, limit, offset, orderBy, sort string) ([]models.JobWithTotalCount, error) {
	return s.repo.GetJobs(search, limit, offset, orderBy, sort)
}

func (s *JobsService) GetJob(id string) (models.Job, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.Job{}, internal.ErrInvalid
	}

	return s.repo.GetJob(idInt)
}

func (s *JobsService) DeleteJob(id string) (models.Job, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.Job{}, internal.ErrInvalid
	}

	return s.repo.DeleteJob(idInt)
}

func (s *JobsService) GetCities(search, limit, offset, orderBy, sort string) ([]models.CitiesWithTotalCount, error) {
	return s.repo.GetCities(search, limit, offset, orderBy, sort)
}
