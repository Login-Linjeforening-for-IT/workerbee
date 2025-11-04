package repositories

import (
	"workerbee/db"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type Audiencerepository interface {
	CreateAudience(audience models.Audience) (models.Audience, error)
	GetAudience(id string) (models.Audience, error)
	GetAudiences(limit, offset int, search, orderBy, sort string) ([]models.AudienceWithTotalCount, error)
	UpdateAudience(audience models.Audience) (models.Audience, error)
	DeleteAudience(id string) (int, error)
}

type audiencerepository struct {
	db *sqlx.DB
}

func NewAudiencerepository(db *sqlx.DB) Audiencerepository {
	return &audiencerepository{db: db}
}

func (r *audiencerepository) CreateAudience(audience models.Audience) (models.Audience, error) {
	return db.AddOneRow(
		r.db,
		"./db/audiences/post_audience.sql",
		audience,
	)
}

func (r *audiencerepository) GetAudience(id string) (models.Audience, error) {
	return db.ExecuteOneRow[models.Audience](
		r.db,
		"./db/audiences/get_audience.sql",
		id,
	)
}

func (r *audiencerepository) GetAudiences(limit, offset int, search, orderBy, sort string) ([]models.AudienceWithTotalCount, error) {
	audiences, err := db.FetchAllElements[models.AudienceWithTotalCount](
		r.db,
		"./db/audiences/get_audiences.sql",
		orderBy, sort, limit, offset, search,
	)

	if err != nil {
		return nil, err
	}

	return audiences, nil
}

func (r *audiencerepository) UpdateAudience(audience models.Audience) (models.Audience, error) {
	return db.AddOneRow(
		r.db,
		"./db/audiences/put_audience.sql",
		audience,
	)
}

func (r *audiencerepository) DeleteAudience(id string) (int, error) {
	return db.ExecuteOneRow[int](
		r.db,
		"./db/audiences/delete_audience.sql",
		id,
	)
}
