package services

import (
	"strconv"
	"strings"
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

var allowedSortColumnsCities = map[string]string{
	"id":   "c.id",
	"name": "c.name",
}

type JobsService struct {
	repo repositories.Jobsrepositories
}

func NewJobsService(repo repositories.Jobsrepositories) *JobsService {
	return &JobsService{repo: repo}
}

func (s *JobsService) CreateJob(job models.NewJob) (models.NewJob, error) {
	newJob, err := s.repo.CreateJob(job)
	return newJob, err
}

func (s *JobsService) GetProtectedJobs(search, limit_str, offset_str, orderBy, sort, jobTypes, skills, cities string) ([]models.JobWithTotalCount, error) {
	orderBySanitized, sortSanitized, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsJobs)
	if ok != nil {
		return nil, internal.ErrInvalid
	}

	jobTypesSlice, err := parseFromStringToSlice(jobTypes)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	skillsSlice, err := parseFromStringToSlice(skills)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	citiesSlice, err := parseFromStringToSlice(cities)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetProtectedJobs(limit, offset, search, orderBySanitized, strings.ToUpper(sortSanitized), jobTypesSlice, skillsSlice, citiesSlice)
}

func (s *JobsService) GetJobs(search, limit_str, offset_str, orderBy, sort, jobTypes, skills, cities string) ([]models.JobWithTotalCount, error) {

	orderBySanitized, sortSanitized, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsJobs)
	if ok != nil {
		return nil, internal.ErrInvalid
	}

	jobTypesSlice, err := parseFromStringToSlice(jobTypes)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	skillsSlice, err := parseFromStringToSlice(skills)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	citiesSlice, err := parseFromStringToSlice(cities)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetJobs(limit, offset, search, orderBySanitized, strings.ToUpper(sortSanitized), jobTypesSlice, skillsSlice, citiesSlice)
}

func (s *JobsService) GetJob(id string) (models.Job, error) {
	return s.repo.GetJob(id)
}

func (s *JobsService) GetJobProtected(id string) (models.Job, error) {
	return s.repo.GetJobProtected(id)
}

func (s *JobsService) GetJobsCities() ([]models.Cities, error) {
	return s.repo.GetJobsCities()
}

func (s *JobsService) GetJobTypes() ([]models.JobType, error) {
	return s.repo.GetJobTypes()
}

func (s *JobsService) GetJobSkills() ([]models.JobSkills, error) {
	return s.repo.GetJobSkills()
}

func (s *JobsService) GetAllJobTypes() ([]string, []string, error) {
	jobTypesEN, jobTypesNO, err := s.repo.GetAllJobTypes()
	if err != nil {
		return nil, nil, err
	}
	return internal.ParsePgArray(jobTypesEN), internal.ParsePgArray(jobTypesNO), nil
}

func (s *JobsService) UpdateJob(id_str string, job models.NewJob) (models.NewJob, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.NewJob{}, internal.ErrInvalid
	}

	job.ID = &id

	return s.repo.UpdateJob(job)
}

func (s *JobsService) DeleteJob(id string) (int, error) {
	return s.repo.DeleteJob(id)
}

func (s *JobsService) GetCities(search, limit_str, offset_str, orderBy, sort string) ([]models.CitiesWithTotalCount, error) {
	orderBySanitized, sortSanitized, err := internal.SanitizeSort(orderBy, sort, allowedSortColumnsCities)
	if err != nil {
		return nil, internal.ErrInvalid
	}
	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}
	return s.repo.GetCities(limit, offset, search, orderBySanitized, strings.ToUpper(sortSanitized))
}

func parseFromStringToSlice(input string) ([]string, error) {
	if input != "" {
		slice, err := internal.ParseCSVToSlice[string](input)
		if err != nil {
			return nil, internal.ErrInvalid
		}
		for i := range slice {
			slice[i] = strings.ToLower(slice[i])
		}
		return slice, nil
	} else {
		return make([]string, 0), nil
	}
}
