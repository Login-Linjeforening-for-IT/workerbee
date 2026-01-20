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

var allowedSortColumnsTypes = map[string]string{
	"id":      "jt.id",
	"name_no": "jt.name_no",
	"name_en": "jt.name_en",
}

type JobsService struct {
	repo repositories.Jobsrepositories
}

func NewJobsService(repo repositories.Jobsrepositories) *JobsService {
	return &JobsService{repo: repo}
}

func (s *JobsService) CreateJob(job models.NewJob) (models.NewJob, error) {
	var err error
	job.Cities, err = parseCitiesAndSkills(job.Cities)
	if err != nil {
		return models.NewJob{}, internal.ErrInvalid
	}

	newJob, err := s.repo.CreateJob(job)
	return newJob, err
}

func (s *JobsService) GetProtectedJobs(search, limit_str, offset_str, orderBy, sort, jobTypes, skills, cities, historical_str string) ([]models.JobWithTotalCount, error) {
	orderBySanitized, sortSanitized, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsJobs)
	if ok != nil {
		return nil, internal.ErrInvalid
	}

	jobTypesSlice, err := internal.ParseFromStringToSlice[int](jobTypes)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	skillsSlice, err := internal.ParseFromStringToSlice[string](skills)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	citiesSlice, err := internal.ParseFromStringToSlice[string](cities)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	historical, err := strconv.ParseBool(historical_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetProtectedJobs(limit, offset, search, orderBySanitized, strings.ToUpper(sortSanitized), jobTypesSlice, skillsSlice, citiesSlice, historical)
}

func (s *JobsService) GetJobs(search, limit_str, offset_str, orderBy, sort, jobTypes, skills, cities string) ([]models.JobWithTotalCount, int, error) {

	orderBySanitized, sortSanitized, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsJobs)
	if ok != nil {
		return nil, 0, internal.ErrInvalid
	}

	jobTypesSlice, err := internal.ParseFromStringToSlice[int](jobTypes)
	if err != nil {
		return nil, 0, internal.ErrInvalid
	}

	skillsSlice, err := internal.ParseFromStringToSlice[string](skills)
	if err != nil {
		return nil, 0, internal.ErrInvalid
	}

	citiesSlice, err := internal.ParseFromStringToSlice[string](cities)
	if err != nil {
		return nil, 0, internal.ErrInvalid
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, 0, internal.ErrInvalid
	}

	jobs, err := s.repo.GetJobs(limit, offset, search, orderBySanitized, strings.ToUpper(sortSanitized), jobTypesSlice, skillsSlice, citiesSlice)
	if err != nil {
		return nil, 0, err
	}

	cacheTTL := s.GetNextPublishTime()

	return jobs, cacheTTL, nil
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

func (s *JobsService) UpdateJob(id_str string, job models.NewJob) (models.NewJob, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.NewJob{}, internal.ErrInvalid
	}

	job.Cities, err = parseCitiesAndSkills(job.Cities)
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

func (s *JobsService) GetAllJobTypes(search, limit_str, offset_str, orderBy, sort string) ([]models.JobTypeWithTotalCount, error) {
	orderBySanitized, sortSanitized, err := internal.SanitizeSort(orderBy, sort, allowedSortColumnsTypes)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	jobTypes, err := s.repo.GetAllJobTypes(limit, offset, search, orderBySanitized, strings.ToUpper(sortSanitized))
	if err != nil {
		return nil, err
	}

	return jobTypes, nil
}

func (s *JobsService) GetOneJobType(id string) (models.JobType, error) {
	return s.repo.GetOneJobType(id)
}

func (s *JobsService) CreateJobType(jobType models.JobType) (models.JobType, error) {
	return s.repo.CreateJobType(jobType)
}

func (s *JobsService) UpdateJobType(id_str string, jobType models.JobType) (models.JobType, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.JobType{}, internal.ErrInvalid
	}

	jobType.ID = id

	return s.repo.UpdateJobType(jobType)
}

func (s *JobsService) DeleteJobType(id string) (int, error) {
	return s.repo.DeleteJobType(id)
}

func (s *JobsService) GetNextPublishTime() int {
	cacheTTL, err := s.repo.GetNextPublishTime()
	if err != nil || cacheTTL == nil {
		return 3600
	}
	return internal.ParseCacheControlHeader(cacheTTL)
}

func parseCitiesAndSkills(cities []string) ([]string, error) {
	for i := range cities {
		cities[i] = internal.FormatNameWithCapitalFirstLetter(cities[i])
	}
	return cities, nil
}
