package repositories

import (
	"workerbee/db"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type Categoryrepository interface {
	GetCategories(limit, offset int, search, orderBy, sort string) ([]models.CategoryWithTotalCount, error)
	GetCategory(id string) (models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	DeleteCategory(id string) (int, error)
}

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) Categoryrepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) CreateCategory(category models.Category) (models.Category, error) {
	return db.AddOneRow(
		r.db,
		"./db/categories/post_category.sql",
		category,
	)
}

func (r *categoryRepository) GetCategories(limit, offset int, search, orderBy, sort string) ([]models.CategoryWithTotalCount, error) {
	categories, err := db.FetchAllElements[models.CategoryWithTotalCount](
		r.db,
		"./db/categories/get_categories.sql",
		orderBy, sort, limit, offset, search,
	)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepository) GetCategory(id string) (models.Category, error) {
	return db.ExecuteOneRow[models.Category](
		r.db,
		"./db/categories/get_category.sql",
		id,
	)
}
func (r *categoryRepository) UpdateCategory(category models.Category) (models.Category, error) {
	return db.AddOneRow(
		r.db,
		"./db/categories/put_category.sql",
		category,
	)
}

func (r *categoryRepository) DeleteCategory(id string) (int, error) {
	catID, err := db.ExecuteOneRow[int](
		r.db,
		"./db/categories/delete_category.sql",
		id,
	)
	if err != nil {
		return 0, err
	}
	return catID, nil
}
