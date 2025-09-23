package handlers

import (
	"workerbee/services"
)

type Handler struct {
	Events      services.EventService
	Rules       services.RuleService
	Stats       services.StatsService
	Forms       services.FormService
	Questions   services.QuestionService
	Submissions services.SubmissionService
	Jobs        services.JobsService
}
