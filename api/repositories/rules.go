package repositories

import (
	"workerbee/db"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type Rulerepositories interface {
	GetRules(search, limit, offset, orderBy, sort string) ([]models.RuleWithTotalCount, error)
}

type rulerepositories struct {
	db *sqlx.DB
}

func NewRulerepositories(db *sqlx.DB) Rulerepositories {
	return &rulerepositories{db: db}
}

func (r *rulerepositories) GetRules(search, limit, offset, orderBy, sort string) ([]models.RuleWithTotalCount, error) {
	rules, err := db.FetchAllElements[models.RuleWithTotalCount](
		r.db,
		"./db/rules/get_rules.sql",
		orderBy, sort,
		limit,
		offset,
		search,
	)
	if err != nil {
		return nil, err
	}
	return rules, nil
}
