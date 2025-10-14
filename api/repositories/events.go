package repositories

import (
	"os"
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Eventrepositories interface {
	GetEvents(limit, offset int, search, orderBy, sort string, historical bool, categories []int) ([]models.EventWithTotalCount, error)
	GetEvent(id string) (models.Event, error)
	GetEventCategories() ([]models.EventCategory, error)
	DeleteEvent(id string) (int, error)
	UpdateOneEvent(id int, event models.NewEvent) (models.NewEvent, error)
	CreateEvent(event models.NewEvent) (models.NewEvent, error)
}

type eventRepositories struct {
	db *sqlx.DB
}

func NewEventrepositories(db *sqlx.DB) Eventrepositories {
	return &eventRepositories{db: db}
}

func (r *eventRepositories) CreateEvent(event models.NewEvent) (models.NewEvent, error) {
	return db.AddOneRow(
		r.db,
		"./db/events/post_event.sql",
		event,
	)
}

func (r *eventRepositories) GetEvents(limit, offset int, search, orderBy, sort string, historical bool, categories []int) ([]models.EventWithTotalCount, error) {
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

func (r *eventRepositories) GetEventCategories() ([]models.EventCategory, error) {
	var categories []models.EventCategory

	sqlBytes, err := os.ReadFile("./db/events/get_categories.sql")
	if err != nil {
		return nil, err
	}

	err = r.db.Select(&categories, string(sqlBytes))
	if err != nil {
		return nil, internal.ErrInvalid
	}
	return categories, nil
}

func (r *eventRepositories) GetEvent(id string) (models.Event, error) {
	event, err := db.ExecuteOneRow[models.Event](r.db, "./db/events/get_event.sql", id)
	if err != nil {
		return models.Event{}, internal.ErrInvalid
	}
	return event, nil
}

func (r *eventRepositories) UpdateOneEvent(id int, event models.NewEvent) (models.NewEvent, error) {
	event.ID = id

	return db.AddOneRow(r.db, "./db/events/put_event.sql", event)
}

func (r *eventRepositories) DeleteEvent(id string) (int, error) {
	eventId, err := db.ExecuteOneRow[int](r.db, "./db/events/delete_event.sql", id)
	if err != nil {
		return 0, internal.ErrInvalid
	}
	return eventId, nil
}
