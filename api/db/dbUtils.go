package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func FetchAllElements[T any](
	db *sqlx.DB,
	sqlPath string,
	orderBy, sort string,
	limit, offset string,
	args ...any,
) ([]T, error) {
	var result []T

	sqlBytes, err := os.ReadFile(sqlPath)
	if err != nil {
		return nil, err
	}
	// append limit and offset to args
	args = append(args, limit, offset)

	query := fmt.Sprintf("%s ORDER BY %s %s \nLIMIT $%d OFFSET $%d;",
		string(sqlBytes),
		orderBy,
		sort,
		len(args)-1, // limit is second to last
		len(args),   // offset is last
	)

	err = db.Select(&result, query, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ExecuteOneRow[T any](db *sqlx.DB, sqlPath, id string) (T, error) {
	var result T

	sqlBytes, err := os.ReadFile(sqlPath)
	if err != nil {
		return result, err
	}

	err = db.Get(&result, string(sqlBytes), id)
	if err != nil {
		return result, err
	}

	return result, nil
}