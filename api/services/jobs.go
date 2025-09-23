package services

import (
	"strconv"
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repositories"
)

var allowedSortColumnsJobs = map[string]string{
	"id":           "ja.id",
	"visible":      "ja.visible",
	"highlight":    "ja.highlight",
	"title_no":     "ja.title_no",
	"title_en":     "ja.title_en",
	"job_type":     "ja.job_type",
	"time_expire":  "ja.time_expire",
	"time_publish": "ja.time_publish",
	"created_at":   "ja.created_at",
	"updated_at":   "ja.updated_at",
}

type JobsService struct {
	repo repositories.Jobsrepositories
}

func NewJobsService(repo repositories.Jobsrepositories) *JobsService {
	return &JobsService{repo: repo}
}

func (s *JobsService) GetJobs(search, limit, offset, orderBy, sort string) ([]models.JobWithTotalCount, error) {
	orderBySanitized, sortSanitized, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsJobs)
	if ok != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetJobs(search, limit, offset, orderBySanitized, sortSanitized)
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
