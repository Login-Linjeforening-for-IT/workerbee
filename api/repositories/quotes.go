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
	UpdateQuote(quote models.BaseQuote) (models.BaseQuote, error)
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
		"./db/quotes/get_quote.sql",
		id,
	)
}

func (r *quoteRepository) UpdateQuote(quote models.BaseQuote) (models.BaseQuote, error) {
	return db.AddOneRow(
		r.db,
		"./db/quotes/put_quote.sql",
		quote,
		
	)
}

func (r *quoteRepository) DeleteQuote(id string) (int, error) {
	deletedID, err := db.DeleteOneRow[int](
		r.db,
		"./db/quotes/delete_quote.sql",
		id,
	)
	return deletedID, err
}