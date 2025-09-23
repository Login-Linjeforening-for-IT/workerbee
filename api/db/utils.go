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

	query := fmt.Sprintf("%s ORDER BY %s %s LIMIT $%d OFFSET $%d",
		string(sqlBytes),
		sort,
		orderBy,
		len(args)-1, // limit is second to last
		len(args),   // offset is last
	)

	err = db.Select(&result, query, args...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
