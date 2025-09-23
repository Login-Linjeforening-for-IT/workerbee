package repositories

import (
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type Rulerepositories interface {
	GetRules(search, limit, offset, orderBy, sort string) ([]models.RuleWithTotalCount, error)
	GetRule(id string) (models.Rule, error)
	DeleteRule(id string) (models.Rule, error)
}

type ruleRepositories struct {
	db *sqlx.DB
}

func NewRulerepositories(db *sqlx.DB) Rulerepositories {
	return &ruleRepositories{db: db}
}

func (r *ruleRepositories) GetRules(search, limit, offset, orderBy, sort string) ([]models.RuleWithTotalCount, error) {
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

func (r *ruleRepositories) GetRule(id string) (models.Rule, error) {
	rule, err := db.ExecuteOneRow[models.Rule](r.db, "./db/rules/get_rule.sql", id)
	if err != nil {
		return rule, internal.ErrInvalid
	}
	return rule, nil
}

func (r *ruleRepositories) DeleteRule(id string) (models.Rule, error) {
	rule, err := db.ExecuteOneRow[models.Rule](r.db, "./db/rules/delete_rule.sql", id)
	if err != nil {
		return rule, internal.ErrInvalid
	}
	return rule, nil
}
