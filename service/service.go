package service

import (
	"context"

	db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"
)

type Service interface {
	db.Store
	GetEventDetails(context.Context, int32) (EventDetails, error)
}

type service struct {
	db.Store
}

func NewService(store db.Store) Service {
	return &service{
		Store: store,
	}
}

type EventDetails struct {
	Event         db.Event                        `json:"event"`
	Category      db.Category                     `json:"category"`
	Rule          *db.Rule                        `json:"rule,omitempty"`
	Location      *db.Location                    `json:"location,omitempty"`
	Organizations []db.GetOrganizationsOfEventRow `json:"organizations"`
	Audiences     []db.Audience                   `json:"audiences"`
}
