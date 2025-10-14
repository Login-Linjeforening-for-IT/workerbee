package repositories

import (
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type OrganizationRepository interface {
	CreateOrg(org models.Organization) (models.Organization, error)
	GetOrgs(limit, offset int, search, orderBy, sort string) ([]models.OrganizationWithTotalCount, error)
	GetOrg(id string) (models.Organization, error)
	UpdateOrg(org models.Organization) (models.Organization, error)
	DeleteOrg(id string) (int, error)
}

type organizationRepository struct {
	db *sqlx.DB
}

func NewOrganizationRepository(db *sqlx.DB) OrganizationRepository {
	return &organizationRepository{db: db}
}

func (r *organizationRepository) GetOrgs(limit, offset int, search, orderBy, sort string) ([]models.OrganizationWithTotalCount, error) {
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

func (r *organizationRepository) DeleteOrg(id string) (int, error) {
	orgId, err := db.ExecuteOneRow[int](
		r.db, "./db/organizations/delete_organization.sql", id,
	)
	if err != nil {
		return 0, internal.ErrInvalid
	}

	return orgId, nil
}

func (r *organizationRepository) UpdateOrg(org models.Organization) (models.Organization, error) {
	return db.AddOneRow(
		r.db,
		"./db/organizations/put_organization.sql",
		org,
	)
}

func (r *organizationRepository) CreateOrg(org models.Organization) (models.Organization, error) {
	return db.AddOneRow(
		r.db,
		"./db/organizations/post_organization.sql",
		org,
	)
}
