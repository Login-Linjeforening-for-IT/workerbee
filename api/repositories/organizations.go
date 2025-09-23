package repositories

import "github.com/jmoiron/sqlx"

type OrganizationRepository interface {
}

type organizationRepository struct {
	db *sqlx.DB
}

func NewOrganizationRepository(db *sqlx.DB) OrganizationRepository {
	return &organizationRepository{db: db}
}
