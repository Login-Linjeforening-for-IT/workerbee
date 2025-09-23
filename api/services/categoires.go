package services

import "workerbee/repositories"

type CategorieService struct {
	repo repositories.CategoireRepository
}

func NewCategorieService(repo repositories.CategoireRepository) *CategorieService {
	return &CategorieService{repo: repo}
}
