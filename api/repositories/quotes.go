package repositories

import (
	"workerbee/db"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type QuoteRepository interface {
	GetQuote(id string) (models.BaseQuote, error)
	CreateQuote(quote models.BaseQuote) (models.BaseQuote, error)
	GetQuotes(limit, offset int) ([]models.QuoteWithTotalCount, error)
	DeleteQuote(id string) (int, error)
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

func (r *quoteRepository) GetQuote(id string) (models.BaseQuote, error) {
	return db.ExecuteOneRow[models.BaseQuote](
		r.db,
		"./db/quotes/get_quote_by_id.sql",
		id,
	)
}

func (r *quoteRepository) DeleteQuote(id string) (int, error) {
	// Here you can implement the actual deletion logic, e.g., executing a DELETE SQL statement.
	return 1, nil
}