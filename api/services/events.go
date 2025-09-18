package services

import (
	"strconv"

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
		return nil, err
	}

	return s.repo.GetEvents(search, limit, offset, orderBy, sort, historicalBool)
}

func (s *EventService) GetEvent(id string) (models.Event, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.Event{}, err
	}

	return s.repo.GetEvent(idInt)
}
