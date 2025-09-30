package services

import (
	"strconv"
	"strings"
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repositories"
)

var allowedSortColumnsEvents = map[string]string{
	"id":           "e.id",
	"name_no":      "e.name_no",
	"name_en":      "e.name_en",
	"time_start":   "e.time_start",
	"time_end":     "e.time_end",
	"time_publish": "e.time_publish",
	"canceled":     "e.canceled",
	"capacity":     "e.capacity",
	"full":         "e.full",
}

type EventService struct {
	repo repositories.Eventrepositories
}

func NewEventService(repo repositories.Eventrepositories) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(body models.Event) (models.Event, error) {
	return s.repo.CreateEvent(body)
}

func (s *EventService) GetEvents(search, limit, offset, orderBy, sort, historical string) ([]models.EventWithTotalCount, error) {
	sanitizedOrderBy, sanitizedSort, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsEvents)
	if ok != nil {
		return nil, internal.ErrInvalid
	}

	historicalBool, err := strconv.ParseBool(historical)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetEvents(search, limit, offset, sanitizedOrderBy, strings.ToUpper(sanitizedSort), historicalBool)
}

func (s *EventService) GetEvent(id string) (models.Event, error) {
	return s.repo.GetEvent(id)
}

func (s *EventService) DeleteEvent(id string) (models.Event, error) {
	return s.repo.DeleteEvent(id)
}
