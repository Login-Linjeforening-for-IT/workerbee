package services

import (
	"slices"
	"strconv"
	"workerbee/repositories"

	ics "github.com/arran4/golang-ical"
)

type CalendarService struct {
	repo repositories.CalendarRepository
}

func NewCalendarService(repo repositories.CalendarRepository) *CalendarService {
	return &CalendarService{repo: repo}
}

func (s *CalendarService) GetCalendarData(categories_str, language string) (*ics.Calendar, error) {
	cal := ics.NewCalendar()
    cal.SetMethod(ics.MethodPublish) 
    cal.SetProductId("-//Login//Events//EN")
    cal.SetName("Login Events")
    cal.SetCalscale("GREGORIAN")

	if !slices.Contains(validLanguages, language) {
		language = "en"
	}

	categories, err := parseToArray(categories_str)
	if err != nil {
		return nil, err
	}

	events, err := s.repo.GetCalendarData(categories)
	if err != nil {
		return nil, err
	}

	for _, event := range events {
		uid := strconv.Itoa(event.ID) + "@login.no"
		calEvent := cal.AddEvent(uid)
		switch language {
		case "no":
			calEvent.SetSummary(event.NameNo)
			calEvent.SetDescription(event.DescriptionNo)
		case "en":
			calEvent.SetSummary(event.NameEn)
			calEvent.SetDescription(event.DescriptionEn)
		}
		calEvent.SetStartAt(event.TimeStart.Time)
		calEvent.SetEndAt(event.TimeEnd.Time)
	}

	return cal, nil
}
