package repositories

import (
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type CategoireRepository interface {
	GetCategories(limit, offset int, search, orderBy, sort string) ([]models.CategoryWithTotalCount, error)
	GetCategory(id string) (models.Category, error)
	DeleteCategory(id string) (models.Category, error)
	CreateCateory(category models.Category) (models.Category, error)
	UpdateCateory(category models.Category, id int) (models.Category, error)
}

type categoireRepository struct {
	db *sqlx.DB
}

func NewCategoireRepository(db *sqlx.DB) CategoireRepository {
	return &categoireRepository{db: db}
}

func (r *categoireRepository) GetCategories(limit, offset int, search, orderBy, sort string) ([]models.CategoryWithTotalCount, error) {
	categories, err := db.FetchAllElements[models.CategoryWithTotalCount](
		r.db,
		"./db/categories/get_categories.sql",
		orderBy, sort,
		limit,
		offset,
		search,
	)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoireRepository) GetCategory(id string) (models.Category, error) {
	category, err := db.ExecuteOneRow[models.Category](r.db, "./db/categories/get_category.sql", id)
	if err != nil {
		return models.Category{}, internal.ErrInvalid
	}
	return category, nil
}

func (r *categoireRepository) DeleteCategory(id string) (models.Category, error) {
	category, err := db.ExecuteOneRow[models.Category](r.db, "./db/categories/delete_category.sql", id)
	if err != nil {
		return models.Category{}, internal.ErrInvalid
	}
	return category, nil
}

func (r *categoireRepository) UpdateCateory(category models.Category, id int) (models.Category, error) {
	return db.AddOneRow(
		r.db,
		"./db/categories/put_category.sql",
		category,
	)
}

func (r *categoireRepository) CreateCateory(category models.Category) (models.Category, error) {
	return db.AddOneRow(
		r.db,
		"./db/categories/post_category.sql",
		category,
	)
}
