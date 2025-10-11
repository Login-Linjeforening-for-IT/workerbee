package repositories

import (
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Eventrepositories interface {
	GetEvents(search, limit, offset, orderBy, sort string, historical bool, categories []int) ([]models.EventWithTotalCount, error)
	GetEvent(id string) (models.Event, error)
	DeleteEvent(id string) (models.Event, error)
	UpdateOneEvent(id int, event models.Event) (models.Event, error)
	CreateEvent(event models.Event) (models.Event, error)
}

type eventRepositories struct {
	db *sqlx.DB
}

func NewEventrepositories(db *sqlx.DB) Eventrepositories {
	return &eventRepositories{db: db}
}

func (r *eventRepositories) CreateEvent(event models.Event) (models.Event, error) {
	return db.AddOneRow(
		r.db,
		"./db/events/post_event.sql",
		event,
	)
}

func (r *eventRepositories) GetEvents(search, limit, offset, orderBy, sort string, historical bool, categories []int) ([]models.EventWithTotalCount, error) {
	events, err := db.FetchAllElements[models.EventWithTotalCount](
		r.db,
		"./db/events/get_events.sql",
		orderBy, sort,
		limit,
		offset,
		search,
		historical,
		pq.Array(categories),
	)

	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepositories) GetEvent(id string) (models.Event, error) {
	event, err := db.ExecuteOneRow[models.Event](r.db, "./db/events/get_event.sql", id)
	if err != nil {
		return models.Event{}, internal.ErrInvalid
	}
	return event, nil
}

func (r *eventRepositories) UpdateOneEvent(id int, event models.Event) (models.Event, error) {
	event.ID = id

	return db.AddOneRow(r.db, "./db/events/put_event.sql", event)
}

func (r *eventRepositories) DeleteEvent(id string) (models.Event, error) {
	event, err := db.ExecuteOneRow[models.Event](r.db, "./db/events/delete_event.sql", id)
	if err != nil {
		return models.Event{}, internal.ErrInvalid
	}
	return event, nil
}
