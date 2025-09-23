package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type JobsRepository interface {
	GetJobs(search, limit, offset, orderBy, sort string) ([]models.JobWithTotalCount, error)
	GetJob(id int) (models.Job, error)
	DeleteJob(id int) (models.Job, error)
	GetCities(search, limit, offset, orderBy, sort string) ([]models.CitiesWithTotalCount, error)
}

type jobsRepository struct {
	db *sqlx.DB
}

func NewJobRepository(db *sqlx.DB) JobsRepository {
	return &jobsRepository{db: db}
}

func (r *jobsRepository) GetJobs(search, limit, offset, orderBy, sort string) ([]models.JobWithTotalCount, error) {
	var jobs []models.JobWithTotalCount

	sqlBytes, err := os.ReadFile("./db/jobs/get_jobs.sql")
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("%s ORDER BY %s %s\nLIMIT $2 OFFSET $3;", string(sqlBytes), sort, orderBy)

	err = r.db.Select(&jobs, query, search, limit, offset)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (r *jobsRepository) GetJob(id int) (models.Job, error) {
	var job models.Job

	sqlBytes, err := os.ReadFile("./db/jobs/get_job.sql")
	if err != nil {
		return job, err
	}

	err = r.db.Get(&job, string(sqlBytes), id)
	if err != nil {
		return job, internal.ErrInvalid
	}

	return job, nil
}

func (r *jobsRepository) DeleteJob(id int) (models.Job, error) {
	var job models.Job

	sqlBytes, err := os.ReadFile("./db/jobs/delete_job.sql")
	if err != nil {
		return job, err
	}

	err = r.db.Get(&job, string(sqlBytes), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return job, internal.ErrInvalid
		}
		return job, err
	}

	return job, nil
}

func (r *jobsRepository) GetCities(search, limit, offset, orderBy, sort string) ([]models.CitiesWithTotalCount, error) {
	var cities []models.CitiesWithTotalCount

	sqlBytes, err := os.ReadFile("./db/jobs/get_cities.sql")
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("%s ORDER BY %s %s\nLIMIT $2 OFFSET $3;", string(sqlBytes), sort, orderBy)

	err = r.db.Select(&cities, query, search, limit, offset)
	if err != nil {
		return nil, err
	}
	return cities, nil
}