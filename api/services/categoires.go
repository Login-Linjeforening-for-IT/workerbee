package services

import (
	"strconv"
	"strings"
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repositories"
)

var allowedSortColumnsCategories = map[string]string{
	"id":         "c.id",
	"color":      "c.color",
	"name_no":    "c.name_no",
	"name_en":    "c.name_en",
	"created_at": "c.created_a",
	"updated_at": "c.updated_at",
}

type CategorieService struct {
	repo repositories.CategoireRepository
}

func NewCategorieService(repo repositories.CategoireRepository) *CategorieService {
	return &CategorieService{repo: repo}
}

func (s *CategorieService) GetCategories(search, limit_str, offset_str, orderBy, sort string) ([]models.CategoryWithTotalCount, error) {
	sanitizedOrderBy, sanitizedSort, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsCategories)
	if ok != nil {
		return nil, internal.ErrInvalid
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetCategories(limit, offset, search, sanitizedOrderBy, strings.ToUpper(sanitizedSort))
}

func (s *CategorieService) GetCategory(id string) (models.Category, error) {
	return s.repo.GetCategory(id)
}

func (s *CategorieService) DeleteCategory(id string) (int, error) {
	return s.repo.DeleteCategory(id)
}

func (s *CategorieService) CreateCateory(category models.Category) (models.Category, error) {
	return s.repo.CreateCateory(category)
}

func (s *CategorieService) UpdateCategory(category models.Category, id_str string) (models.Category, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.Category{}, internal.ErrInvalid
	}

	category.ID = id

	return s.repo.UpdateCateory(category, id)
}
