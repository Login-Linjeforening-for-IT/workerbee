package repositories

import (
	"database/sql"
	"errors"
	"os"
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type Eventrepositories interface {
	GetEvents(search, limit, offset, orderBy, sort string, historical bool) ([]models.EventWithTotalCount, error)
	GetEvent(id int) (models.Event, error)
	DeleteEvent(id int) (models.Event, error)
}

type eventRepositories struct {
	db *sqlx.DB
}

func NewEventrepositories(db *sqlx.DB) Eventrepositories {
	return &eventRepositories{db: db}
}

func (r *eventRepositories) GetEvents(search, limit, offset, orderBy, sort string, historical bool) ([]models.EventWithTotalCount, error) {
	events, err := db.FetchAllElements[models.EventWithTotalCount](
		r.db,
		"./db/events/get_events.sql",
		orderBy, sort,
		limit,
		offset,
		search,
		historical,
	)

	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepositories) GetEvent(id int) (models.Event, error) {
	var event models.Event

	sqlBytes, err := os.ReadFile("./db/events/get_event.sql")
	if err != nil {
		return event, err
	}

	err = r.db.Get(&event, string(sqlBytes), id)
	if err != nil {
		return event, internal.ErrInvalid
	}

	return event, nil
}

func (r *eventRepositories) DeleteEvent(id int) (models.Event, error) {
	var event models.Event

	sqlBytes, err := os.ReadFile("./db/events/delete_event.sql")
	if err != nil {
		return event, err
	}

	err = r.db.Get(&event, string(sqlBytes), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return event, internal.ErrInvalid
		}
		return event, err
	}

	return event, nil
}
