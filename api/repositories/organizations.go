package repositories

import (
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type OrganizationRepository interface {
	GetOrgs(search, limit, offset, orderBy, sort string) ([]models.OrganizationWithTotalCount, error)
	GetOrg(id string) (models.Organization, error)
	DeleteOrg(id string) (models.Organization, error)
}

type organizationRepository struct {
	db *sqlx.DB
}

func NewOrganizationRepository(db *sqlx.DB) OrganizationRepository {
	return &organizationRepository{db: db}
}

func (r *organizationRepository) GetOrgs(search, limit, offset, orderBy, sort string) ([]models.OrganizationWithTotalCount, error) {
	orgs, err := db.FetchAllElements[models.OrganizationWithTotalCount](
		r.db, "./db/organizations/get_organizations.sql",
		orderBy, sort, limit, offset, search,
	)
	if err != nil {
		return nil, err
	}

	return orgs, nil
}

func (r *organizationRepository) GetOrg(id string) (models.Organization, error) {
	org, err := db.ExecuteOneRow[models.Organization](
		r.db, "./db/organizations/get_organization.sql", id,
	)
	if err != nil {
		return models.Organization{}, internal.ErrInvalid
	}

	return org, nil
}

func (r *organizationRepository) DeleteOrg(id string) (models.Organization, error) {
	org, err := db.ExecuteOneRow[models.Organization](
		r.db, "./db/organizations/delete_organization.sql", id,
	)
	if err != nil {
		return models.Organization{}, internal.ErrInvalid
	}

	return org, nil
}
