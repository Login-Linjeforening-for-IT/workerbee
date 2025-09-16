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

	events, err := s.repo.GetEvents(search, limit, offset, orderBy, sort, historicalBool)
	if err != nil {
		return nil, err
	}
	return events, nil
}
