package service

import db "git.logntnu.no/tekkom/web/beehive/admin-api/db/sqlc"

type Service interface {
	db.Store
}

type service struct {
	db.Store
}

func NewService(store db.Store) Service {
	return &service{
		Store: store,
	}
}
