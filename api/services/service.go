package services

import (
	"workerbee/repositories"

	"github.com/go-playground/validator/v10"
)

type Services struct {
	Events        *EventService
	Locations     *LocationService
	Organizations *OrganizationService
	Forms         *FormService
	Jobs          *JobsService
	Questions     *QuestionService
	Rules         *RuleService
	Stats         *StatsService
	Submissions   *SubmissionService
	Validate      *validator.Validate
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		Events:        NewEventService(repos.Events),
		Forms:         NewFormService(repos.Forms),
		Jobs:          NewJobsService(repos.Jobs),
		Questions:     NewQuestionService(repos.Questions),
		Rules:         NewRuleService(repos.Rules),
		Stats:         NewStatsService(repos.Stats),
		Submissions:   NewSubmissionService(repos.Submissions),
		Locations:     NewLocationService(repos.Locations),
		Organizations: NewOrganizationService(repos.Organizations),
		Validate:      validator.New(),
	}
}
