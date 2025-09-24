package services

import (
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

func (s *CategorieService) GetCategories(search, limit, offset, orderBy, sort string) ([]models.CategoryWithTotalCount, error) {
	sanitizedOrderBy, sanitizedSort, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsCategories)
	if ok != nil {
		return nil, internal.ErrInvalid
	}
	return s.repo.GetCategories(search, limit, offset, sanitizedOrderBy, strings.ToUpper(sanitizedSort))
}

func (s *CategorieService) GetCategory(id string) (models.Category, error) {
	return s.repo.GetCategory(id)
}

func (s *CategorieService) DeleteCategory(id string) (models.Category, error) {
	return s.repo.DeleteCategory(id)
}
