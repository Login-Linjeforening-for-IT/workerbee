package services

import (
	"strconv"
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repositories"
)

var allowedSortColumnsCategories = map[string]string{
	"id":      "c.id",
	"name_no": "c.name_no",
	"name_en": "c.name_en",
	"color":   "c.color",
}

type CategoryService struct {
	repo repositories.Categoryrepository
}

func NewCategoryService(repo repositories.Categoryrepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category models.Category) (models.Category, error) {
	return s.repo.CreateCategory(category)
}

func (s *CategoryService) GetCategory(id string) (models.Category, error) {
	return s.repo.GetCategory(id)
}

func (s *CategoryService) GetCategories(search, limit_str, offset_str, orderBy, sort string) ([]models.CategoryWithTotalCount, error) {
	orderBySanitized, sortSanitized, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsCategories)
	if ok != nil {
		return nil, internal.ErrInvalid
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetCategories(limit, offset, search, orderBySanitized, sortSanitized)
}

func (s *CategoryService) UpdateCategory(id_str string, category models.Category) (models.Category, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.Category{}, internal.ErrInvalid
	}
	category.ID = &id
	return s.repo.UpdateCategory(category)
}

func (s *CategoryService) DeleteCategory(id string) (int, error) {
	return s.repo.DeleteCategory(id)
}
