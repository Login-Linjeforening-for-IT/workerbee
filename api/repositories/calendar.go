package repositories

import (
	"os"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type CalendarRepository interface {
	GetCalendarData(categories []int) ([]models.CalendarEvent, error)
}

type calendarRepository struct {
	db *sqlx.DB
}

func NewCalendarRepository(db *sqlx.DB) CalendarRepository {
	return &calendarRepository{
		db: db,
	}
}

func (r *calendarRepository) GetCalendarData(categories []int) ([]models.CalendarEvent, error) {
	var events []models.CalendarEvent

	sqlBytes, err := os.ReadFile("./db/calendar/get_calendar_data.sql")
	if err != nil {
		return nil, err
	}

	if err := r.db.Select(&events, string(sqlBytes), pq.Array(categories)); err != nil {
		return nil, err
	}

	return events, nil
}
