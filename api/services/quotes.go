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

func (s *QuoteService) GetQuote(id string) (models.BaseQuote, error) {
	return s.repo.GetQuote(id)
}

func (s *QuoteService) DeleteQuote(id_str, userID string, admin bool) (int, error) {
	quote, err := s.repo.GetQuote(id_str)
	if err != nil {
		return 0, err
	}

	if quote.Author != userID || !admin {
		return 0, internal.ErrUnauthorized
	}

	return s.repo.DeleteQuote(id_str)
}