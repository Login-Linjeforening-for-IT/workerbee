package services

import (
	"workerbee/repositories"
)

type AudienceService struct {
	repo repositories.Audiencerepository
}

func NewAudienceService(repo repositories.Audiencerepository) *AudienceService {
	return &AudienceService{repo: repo}
}
