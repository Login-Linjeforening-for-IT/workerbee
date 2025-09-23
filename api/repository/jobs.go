package repository

import (
	"fmt"
	"os"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type JobsRepository interface {
	GetJobs(search, limit, offset, orderBy, sort string) ([]models.JobWithTotalCount, error)
	GetJob(id int) (models.Job, error)
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
