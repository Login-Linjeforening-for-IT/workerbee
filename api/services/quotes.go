package services

import (
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repositories"
)

type QuoteService struct {
	repo repositories.QuoteRepository
}

func NewQuoteService(repo repositories.QuoteRepository) *QuoteService {
	return &QuoteService{repo: repo}
}

func (s *QuoteService) CreateQuote(quote models.BaseQuote) (models.BaseQuote, error) {
	return s.repo.CreateQuote(quote)
}

func (s *QuoteService) GetQuotes(limit_str, offset_str string) ([]models.QuoteWithTotalCount, error) {
	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetQuotes(limit, offset)
}
