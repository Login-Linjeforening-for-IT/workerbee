package services

import (
	"strconv"

	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/internal"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/repository"
)

type EventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) GetEvents(search, limit, offset, orderBy, sort, historical string) ([]models.EventWithTotalCount, error) {
	historicalBool, err := strconv.ParseBool(historical)
	if err != nil {
		return nil, internal.ErrInvalidSort
	}

	return s.repo.GetEvents(search, limit, offset, orderBy, sort, historicalBool)
}

func (s *EventService) GetEvent(id string) (models.Event, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.Event{}, internal.ErrInvalidId
	}

	return s.repo.GetEvent(idInt)
}

func (s *EventService) DeleteEvent(id string) (models.Event, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.Event{}, internal.ErrInvalidId
	}
	
	return s.repo.DeleteEvent(idInt)
}