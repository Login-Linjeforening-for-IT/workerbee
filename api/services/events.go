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

func (s *EventService) CreateEvent(body models.NewEvent) (models.NewEvent, error) {
	return s.repo.CreateEvent(body)
}

func (s *EventService) UpdateEvent(body models.NewEvent, id_str string) (models.NewEvent, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.NewEvent{}, internal.ErrInvalid
	}

	return s.repo.UpdateOneEvent(id, body)
}

func (s *EventService) GetEvents(search, limit_str, offset_str, orderBy, sort, historical, categories_str string) ([]models.EventWithTotalCount, error) {
	sanitizedOrderBy, sanitizedSort, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsEvents)
	if ok != nil {
		return nil, internal.ErrInvalid
	}

	historicalBool, err := strconv.ParseBool(historical)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	var categories []string
	if categories_str != "" {
		categories, err = internal.ParseCSVToSlice[string](categories_str)
		if err != nil {
			return nil, internal.ErrInvalid
		}
		for i := range categories {
			categories[i] = strings.ToLower(categories[i])
		}

	} else {
		categories = make([]string, 0)
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetEvents(limit, offset, search, sanitizedOrderBy, strings.ToUpper(sanitizedSort), historicalBool, categories)
}

func (s *EventService) GetEvent(id string) (models.Event, error) {
	return s.repo.GetEvent(id)
}

func (s *EventService) GetEventCategories() ([]models.EventCategory, error) {
	return s.repo.GetEventCategories()
}

func (s *EventService) DeleteEvent(id string) (int, error) {
	return s.repo.DeleteEvent(id)
}
