package repositories

import (
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type Jobsrepositories interface {
	GetJobs(search, limit, offset, orderBy, sort string) ([]models.JobWithTotalCount, error)
	GetJob(id string) (models.Job, error)
	DeleteJob(id string) (models.Job, error)
	GetCities(search, limit, offset, orderBy, sort string) ([]models.CitiesWithTotalCount, error)
}

type jobsrepositories struct {
	db *sqlx.DB
}

func NewJobrepositories(db *sqlx.DB) Jobsrepositories {
	return &jobsrepositories{db: db}
}

func (r *jobsrepositories) GetJobs(search, limit, offset, orderBy, sort string) ([]models.JobWithTotalCount, error) {
	jobs, err := db.FetchAllElements[models.JobWithTotalCount](
		r.db,
		"./db/jobs/get_jobs.sql",
		orderBy, sort,
		limit,
		offset,
		search,
	)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (r *jobsrepositories) GetJob(id string) (models.Job, error) {
	job, err := db.ExecuteOneRow[models.Job](r.db, "./db/jobs/get_job.sql", id)
	if err != nil {
		return models.Job{}, internal.ErrInvalid
	}

	return job, nil
}

func (r *jobsrepositories) DeleteJob(id string) (models.Job, error) {
	job, err := db.ExecuteOneRow[models.Job](r.db, "./db/jobs/delete_job.sql", id)
	if err != nil {
		return models.Job{}, internal.ErrInvalid
	}

	return job, nil
}

func (r *jobsrepositories) GetCities(search, limit, offset, orderBy, sort string) ([]models.CitiesWithTotalCount, error) {
	cities, err := db.FetchAllElements[models.CitiesWithTotalCount](
		r.db,
		"./db/jobs/get_cities.sql",
		orderBy, sort,
		limit,
		offset,
		search,
	)
	if err != nil {
		return nil, err
	}
	return cities, nil
}
