package repository

import (
	"os"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type JobsRepository interface {
	GetJobs(search, limit, offset string) ([]models.JobWithTotalCount, error)
}

type jobsRepository struct {
	db *sqlx.DB
}

func NewJobRepository(db *sqlx.DB) JobsRepository {
	return &jobsRepository{db: db}
}

func (r *jobsRepository) GetJobs(search, limit, offset string) ([]models.JobWithTotalCount, error) {
	var jobs []models.JobWithTotalCount

	sqlBytes, err := os.ReadFile("./db/jobs/get_jobs.sql")
	if err != nil {
		return nil, err
	}

	query := string(sqlBytes)

	err = r.db.Select(&jobs, query, search, limit, offset)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}
