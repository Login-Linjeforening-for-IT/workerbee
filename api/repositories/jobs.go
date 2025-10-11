package repositories

import (
	"database/sql"
	"log"
	"os"
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type Jobsrepositories interface {
	CreateJob(job models.Job) error
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

func (r *jobsrepositories) CreateJob(job models.Job) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var skillIDs []int
	for _, skillName := range job.Skills {
		var skillID int

		err = tx.QueryRow(`SELECT id FROM skills WHERE LOWER(name) = LOWER($1)`, skillName).Scan(&skillID)
		log.Println("Skill ID", skillID)
		if err == sql.ErrNoRows {
			err = tx.QueryRow(`
				INSERT INTO skills (name) 
				VALUES ($1) RETURNING id
			`, skillName).Scan(&skillID)
		}

		if err != nil {
			return err
		}
		skillIDs = append(skillIDs, skillID)
	}

	sqlFile, err := os.ReadFile("./db/jobs/post_job.sql")
	if err != nil {
		return err
	}

	row := tx.QueryRow(
		string(sqlFile),
		job.Visible,
		job.Highlight,
		job.TitleNo,
		job.TitleEn,
		job.PositionTitleNo,
		job.PositionTitleEn,
		job.DescriptionShortNo,
		job.DescriptionShortEn,
		job.DescriptionLongNo,
		job.DescriptionLongEn,
		job.JobType,
		job.TimeExpire,
		job.ApplicationDeadline,
		job.BannerImage,
		job.OrganizationID,
		job.ApplicationURL,
	)
	// Example: scan the returned id (adjust as needed)
	var insertedID int
	err = row.Scan(&insertedID)
	if err != nil {
		return err
	}

	for _, skillID := range skillIDs {
		_, err = tx.Exec(`
			INSERT INTO ad_skill_relation (job_id, skill_id) 
			VALUES ($1, $2)
		`, insertedID, skillID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
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
