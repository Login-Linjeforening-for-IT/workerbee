package db

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

// enables use in type switch
type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "not found"
}

type ForeignKeyViolationError struct {
	Table      string
	Constraint string
}

func (e ForeignKeyViolationError) Error() string {
	return fmt.Sprintf("foreign key violation on table %s for constraint %s", e.Table, e.Constraint)
}

type UniqueViolationError struct {
	Table      string
	Constraint string
}

func (e UniqueViolationError) Error() string {
	return fmt.Sprintf("unique violation on table %s for constraint %s", e.Table, e.Constraint)
}

// function that makes a lib/pq error into a more readable custom error
func ParseError(err error) error {
	if err == nil {
		return nil
	}

	switch err {
	case sql.ErrNoRows:
		return &NotFoundError{}
	}

	switch err := err.(type) {
	case *pq.Error:
		return parseError(err)
	}

	return err
}

func parseError(err *pq.Error) error {
	switch err.Code {
	case "23503": // foreign_key_violation
		return &ForeignKeyViolationError{
			Table:      err.Table,
			Constraint: err.Constraint,
		}
	case "23505": // unique_violation
		return &UniqueViolationError{
			Table:      err.Table,
			Constraint: err.Constraint,
		}
	}

	return err
}
