package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Audiencerepository interface {
	// Define methods for the Audience repository here
}

type audiencerepository struct {
	db *sqlx.DB
}

func NewAudiencerepository(db *sqlx.DB) Audiencerepository {
	return &audiencerepository{db: db}
}
