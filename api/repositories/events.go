package repositories

import (
	"log"
	"os"
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Eventrepositories interface {
	GetProtectedEvents(limit, offset int, search, orderBy, sort string, historical bool, categories []int) ([]models.EventWithTotalCount, error)
	GetEvents(limit, offset int, search, orderBy, sort string, historical bool, categories []int) ([]models.EventWithTotalCount, error)
	GetEvent(id string) (models.Event, error)
	GetProtectedEvent(id string) (models.Event, error)
	GetEventCategories() ([]models.EventCategory, error)
	DeleteEvent(id string) (int, error)
	UpdateOneEvent(id int, event models.NewEvent) (models.NewEvent, error)
	CreateEvent(event models.NewEvent) (models.NewEvent, error)
	GetEventAudiences() (string, string, error)
	GetAllTimeTypes() (string, error)
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

func (r *eventRepositories) GetProtectedEvents(limit, offset int, search, orderBy, sort string, historical bool, categories []int) ([]models.EventWithTotalCount, error) {
	events, err := db.FetchAllElements[models.EventWithTotalCount](
		r.db,
		"./db/events/get_protected_events.sql",
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

	log.Println(categories)

	return categories, nil
}

func (r *eventRepositories) GetEvent(id string) (models.Event, error) {
	event, err := db.ExecuteOneRow[models.Event](r.db, "./db/events/get_event.sql", id)
	if err != nil {
		return models.Event{}, internal.ErrInvalid
	}
	return event, nil
}

func (r *eventRepositories) GetProtectedEvent(id string) (models.Event, error) {
	event, err := db.ExecuteOneRow[models.Event](r.db, "./db/events/get_protected_event.sql", id)
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

func (r *eventRepositories) GetEventAudiences() (string, string, error) {
	audiencesEN, err := db.FetchAllEnumTypes(
		r.db,
		"./db/events/get_all_audiences_en.sql",
	)
	if err != nil {
		return "", "", err
	}

	audiencesNO, err := db.FetchAllEnumTypes(
		r.db,
		"./db/events/get_all_audiences_no.sql",
	)

	if err != nil {
		return "", "", err
	}
	return audiencesEN, audiencesNO, nil
}

func (r *eventRepositories) GetAllTimeTypes() (string, error) {
	timeTypes, err := db.FetchAllEnumTypes(
		r.db,
		"./db/events/get_all_time_types.sql",
	)
	if err != nil {
		return "", err
	}
	return timeTypes, nil
}
