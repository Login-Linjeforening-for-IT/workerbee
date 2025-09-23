package services

import (
	"strconv"
	"workerbee/internal"
	"workerbee/models"
	"workerbee/repository"
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
		return nil, internal.ErrInvalid
	}

	return s.repo.GetEvents(search, limit, offset, orderBy, sort, historicalBool)
}

func (s *EventService) GetEvent(id string) (models.Event, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.Event{}, internal.ErrInvalid
	}

	return s.repo.GetEvent(idInt)
}

func (s *EventService) DeleteEvent(id string) (models.Event, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return models.Event{}, internal.ErrInvalid
	}
	
	return s.repo.DeleteEvent(idInt)
}