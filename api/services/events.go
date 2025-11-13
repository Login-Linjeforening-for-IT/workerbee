package services

import (
	"slices"
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
	allowedTimeTypes, err := s.GetAllTimeTypes()
	if err != nil {
		return models.NewEvent{}, err
	}

	if !slices.Contains(allowedTimeTypes, body.TimeType) {
		return models.NewEvent{}, internal.ErrInvalidTimeType
	}

	return s.repo.CreateEvent(body)
}

func (s *EventService) UpdateEvent(body models.NewEvent, id_str string) (models.NewEvent, error) {
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return models.NewEvent{}, internal.ErrInvalid
	}

	allowedTimeTypes, err := s.GetAllTimeTypes()
	if err != nil {
		return models.NewEvent{}, err
	}

	if !slices.Contains(allowedTimeTypes, body.TimeType) {
		return models.NewEvent{}, internal.ErrInvalidTimeType
	}

	return s.repo.UpdateOneEvent(id, body)
}

func (s *EventService) GetAllTimeTypes() ([]string, error) {
	rawString, err := s.repo.GetAllTimeTypes()
	if err != nil {
		return nil, err
	}
	return internal.ParsePgArray(rawString), nil
}

func (s *EventService) GetEventAudiences() ([]models.Audience, error) {
	audiences, err := s.repo.GetEventAudiences()
	if err != nil {
		return nil, err
	}
	return audiences, nil
}

func (s *EventService) GetEvents(search, limit_str, offset_str, orderBy, sort, categories_str, audiences_str string) ([]models.EventWithTotalCount, error) {
	sanitizedOrderBy, sanitizedSort, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsEvents)
	if ok != nil {
		return nil, internal.ErrInvalid
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	categories, err := parseToArray(categories_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	audiences, err := parseToArray(audiences_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetEvents(limit, offset, search, sanitizedOrderBy, strings.ToUpper(sanitizedSort), categories, audiences)
}

func (s *EventService) GetProtectedEvents(search, limit_str, offset_str, orderBy, sort, historical, categories_str, audiences_str string) ([]models.EventWithTotalCount, error) {
	sanitizedOrderBy, sanitizedSort, ok := internal.SanitizeSort(orderBy, sort, allowedSortColumnsEvents)
	if ok != nil {
		return nil, internal.ErrInvalid
	}

	historicalBool, err := strconv.ParseBool(historical)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	offset, limit, err := internal.CalculateOffset(offset_str, limit_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	categories, err := parseToArray(categories_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	audiences, err := parseToArray(audiences_str)
	if err != nil {
		return nil, internal.ErrInvalid
	}

	return s.repo.GetProtectedEvents(limit, offset, search, sanitizedOrderBy, strings.ToUpper(sanitizedSort), historicalBool, categories, audiences)
}

func (s *EventService) GetEvent(id string) (models.Event, error) {
	return s.repo.GetEvent(id)
}

func (s *EventService) GetProtectedEvent(id string) (models.Event, error) {
	return s.repo.GetProtectedEvent(id)
}

func (s *EventService) GetEventCategories() ([]models.EventCategory, error) {
	return s.repo.GetEventCategories()
}

func (s *EventService) DeleteEvent(id string) (int, error) {
	return s.repo.DeleteEvent(id)
}

func (s *EventService) GetEventNames() ([]models.EventName, error) {
	return s.repo.GetEventNames()
}

func parseToArray(content string) ([]int, error) {
	if content != "" {
		categories, err := internal.ParseCSVToSlice[int](content)
		if err != nil {
			return nil, internal.ErrInvalid
		}
		return categories, nil
	} else {
		return make([]int, 0), nil
	}
}
