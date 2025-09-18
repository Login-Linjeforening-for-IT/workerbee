package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/internal"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
)

type EventRepository interface {
	GetEvents(search, limit, offset, orderBy, sort string, historical bool) ([]models.EventWithTotalCount, error)
	GetEvent(id int) (models.Event, error)
}

type eventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) GetEvents(search, limit, offset, orderBy, sort string, historical bool) ([]models.EventWithTotalCount, error) {
	var events []models.EventWithTotalCount

	sqlBytes, err := os.ReadFile("./db/events/get_events.sql")
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("%s ORDER BY %s %s\nLIMIT $2 OFFSET $3;", string(sqlBytes), sort, orderBy)

	err = r.db.Select(&events, query, search, limit, offset, historical)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepository) GetEvent(id int) (models.Event, error) {
	var event models.Event

	sqlBytes, err := os.ReadFile("./db/events/get_event.sql")
	if err != nil {
		return event, err
	}

	err = r.db.Get(&event, string(sqlBytes), id)
	if err != nil {
		return event, internal.ErrInvalidSort
	}

	return event, nil
}
