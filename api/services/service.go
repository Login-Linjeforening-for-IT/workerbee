package services

import "workerbee/repositories"

type Services struct {
	Events      *EventService
	Forms       *FormService
	Jobs        *JobsService
	Questions   *QuestionService
	Rules       *RuleService
	Stats       *StatsService
	Submissions *SubmissionService
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		Events:      NewEventService(repos.Events),
		Forms:       NewFormService(repos.Forms),
		Jobs:        NewJobsService(repos.Jobs),
		Questions:   NewQuestionService(repos.Questions),
		Rules:       NewRuleService(repos.Rules),
		Stats:       NewStatsService(repos.Stats),
		Submissions: NewSubmissionService(repos.Submissions),
	}
}
