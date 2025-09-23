package repositories

import (
	"fmt"
	"os"
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
	var rules []models.RuleWithTotalCount

	sqlBytes, err := os.ReadFile("./db/rules/get_rules.sql")
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("%s ORDER BY %s %s\nLIMIT $2 OFFSET $3;", string(sqlBytes), sort, orderBy)

	err = r.db.Select(&rules, query, search, limit, offset)
	if err != nil {
		return nil, err
	}
	return rules, nil
}
