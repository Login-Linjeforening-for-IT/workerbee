package repository

import (
	"fmt"
	"os"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type RuleRepository interface {
	GetRules(search, limit, offset, orderBy, sort string) ([]models.RuleWithTotalCount, error)
}

type ruleRepository struct {
	db *sqlx.DB
}

func NewRuleRepository(db *sqlx.DB) RuleRepository {
	return &ruleRepository{db: db}
}

func (r *ruleRepository) GetRules(search, limit, offset, orderBy, sort string) ([]models.RuleWithTotalCount, error) {
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
