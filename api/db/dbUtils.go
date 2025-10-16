package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"workerbee/internal"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func FetchAllEnumTypes(db *sqlx.DB, sqlPath string) (string, error) {
	var result string

	sqlBytes, err := os.ReadFile(sqlPath)
	if err != nil {
		return "", err
	}

	err = db.Get(&result, string(sqlBytes))
	if err != nil {
		return "", err
	}

	return result, nil
}

func FetchAllElements[T any](
	db *sqlx.DB,
	sqlPath string,
	orderBy, sort string,
	limit, offset int,
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
		len(args)-1,
		len(args),
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

func AddOneRow[T any](db *sqlx.DB, sqlPath string, body T) (T, error) {
	var pqErr *pq.Error
	var result T

	sqlBytes, err := os.ReadFile(sqlPath)
	if err != nil {
		return result, err
	}

	rows, err := db.NamedQuery(string(sqlBytes), body)
	if err != nil {
		if errors.As(err, &pqErr) && pqErr.Code == "23503" {
			return result, internal.ErrInvalidForeignKey
		}
		log.Println("got error ", err)
		return result, internal.ErrInvalid
	}
	defer rows.Close()

	if !rows.Next() {
		return result, internal.ErrNotFound
	}

	err = rows.StructScan(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
