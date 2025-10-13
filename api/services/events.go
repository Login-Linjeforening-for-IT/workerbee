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

func (s *EventService) UpdateEvent(body models.Event, id_str string) (models.Event, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.Event{}, internal.ErrInvalid
	}

	return s.repo.UpdateOneEvent(id, body)
}

func (s *EventService) GetEvents(search, limit, offset, orderBy, sort, historical, categories_str string) ([]models.EventWithTotalCount, error) {
	sanitizedOrderBy, sanitizedSort, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsEvents)
	if ok != nil {
		return nil, internal.ErrInvalid
	}

	historicalBool, err := strconv.ParseBool(historical)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	var numbers []int
	if categories_str != "" {
		numbers, err = internal.ParseCSVToSlice[int](categories_str)
		if err != nil {
			return nil, internal.ErrInvalid
		}
	} else {
		numbers = make([]int, 0)
	}

	offset, err = internal.CalculateOffset(offset, limit)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetEvents(search, limit, offset, sanitizedOrderBy, strings.ToUpper(sanitizedSort), historicalBool, numbers)
}

func (s *EventService) GetEvent(id string) (models.Event, error) {
	return s.repo.GetEvent(id)
}

func (s *EventService) GetEventCategories() ([]models.EventCategory, error) {
	return s.repo.GetEventCategories()
}

func (s *EventService) DeleteEvent(id string) (models.Event, error) {
	return s.repo.DeleteEvent(id)
}
