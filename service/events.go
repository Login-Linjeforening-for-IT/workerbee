package service

import (
	"context"
)

func (service *service) GetEventDetails(ctx context.Context, id int32) (EventDetails, error) {
	var eventDetails EventDetails

	// * Event
	event, err := service.GetEvent(ctx, id)
	if err != nil {
		return eventDetails, err
	}

	eventDetails.Event = event

	// * Category
	category, err := service.GetCategory(ctx, event.Category)
	if err != nil {
		return eventDetails, err
	}

	eventDetails.Category = category

	// * Rule
	if event.Rule.Valid {
		rule, err := service.GetRule(ctx, int32(event.Rule.Int64))
		if err != nil {
			return eventDetails, err
		}

		eventDetails.Rule = &rule
	}

	// * Location
	if event.Location.Valid {
		location, err := service.GetLocation(ctx, int32(event.Location.Int64))
		if err != nil {
			return eventDetails, err
		}

		eventDetails.Location = &location
	}

	// * Organizations
	organizations, err := service.GetOrganizationsOfEvent(ctx, id)
	if err != nil {
		return eventDetails, err
	}

	eventDetails.Organizations = organizations

	// * Audiences
	audiences, err := service.GetAudiencesOfEvent(ctx, id)
	if err != nil {
		return eventDetails, err
	}

	eventDetails.Audiences = audiences

	return eventDetails, nil
}
