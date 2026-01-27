package repositories

import (
	"workerbee/db"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type QuoteRepository interface {
	CreateQuote(quote models.BaseQuote) (models.BaseQuote, error)
	GetQuotes(limit, offset int) ([]models.QuoteWithTotalCount, error)
}

type quoteRepository struct {
	db *sqlx.DB
}

func NewQuoteRepository(db *sqlx.DB) QuoteRepository {
	return &quoteRepository{
		db: db,
	}
}

func (r *quoteRepository) CreateQuote(quote models.BaseQuote) (models.BaseQuote, error) {
	return db.AddOneRow(
		r.db,
		"./db/quotes/post_quote.sql",
		quote,
	)
}

func (r *quoteRepository) GetQuotes(limit, offset int) ([]models.QuoteWithTotalCount, error) {
	quotes, err := db.SelectWithLimitOffset[models.QuoteWithTotalCount](
		r.db,
		"./db/quotes/get_quotes.sql",
		limit,
		offset,
	)

	if err != nil {
		return nil, err
	}

	return quotes, nil
}
