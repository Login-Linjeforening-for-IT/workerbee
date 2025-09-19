package handlers

import (
	"workerbee/services"
)
type Handler struct {
	Events      services.EventService
	Stats       services.StatsService
	Forms       services.FormService
	Questions   services.QuestionService
	Submissions services.SubmissionService
}