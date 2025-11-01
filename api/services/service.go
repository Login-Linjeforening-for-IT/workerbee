package services

import (
	"workerbee/repositories"

	"github.com/go-playground/validator/v10"
)

type Services struct {
	Audiences     *AudienceService
	Categories    *CategoryService
	Events        *EventService
	Locations     *LocationService
	Organizations *OrganizationService
	Forms         *FormService
	Jobs          *JobsService
	Questions     *QuestionService
	Rules         *RuleService
	Stats         *StatsService
	Submissions   *SubmissionService
	ImageService  *ImageService
	Validate      *validator.Validate
	Honey         *HoneyService
	Alerts        *AlertService
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		Audiences:     NewAudienceService(repos.Audiences),
		Categories:    NewCategoryService(repos.Categories),
		Events:        NewEventService(repos.Events),
		Forms:         NewFormService(repos.Forms),
		Jobs:          NewJobsService(repos.Jobs),
		Questions:     NewQuestionService(repos.Questions),
		Rules:         NewRuleService(repos.Rules),
		Stats:         NewStatsService(repos.Stats),
		Submissions:   NewSubmissionService(repos.Submissions),
		Locations:     NewLocationService(repos.Locations),
		Organizations: NewOrganizationService(repos.Organizations),
		ImageService:  NewImageService(repos.Images),
		Honey:         NewHoneyService(repos.Honey),
		Alerts:        NewAlertService(repos.Alerts),
		Validate:      validator.New(),
	}
}
