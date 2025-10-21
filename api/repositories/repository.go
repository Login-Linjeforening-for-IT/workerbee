package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Audiences     Audiencerepository
	Categories    Categoryrepository
	Forms         Formrepositories
	Locations     LocationRepository
	Organizations OrganizationRepository
	Events        Eventrepositories
	Jobs          Jobsrepositories
	Questions     Questionrepositories
	Rules         Rulerepositories
	Stats         Statsrepositories
	Submissions   Submissionrepositories
	Honey         HoneyRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Audiences:     NewAudiencerepository(db),
		Categories:    NewCategoryRepository(db),
		Forms:         NewFormrepositories(db),
		Events:        NewEventrepositories(db),
		Jobs:          NewJobrepositories(db),
		Questions:     NewQuestionrepositories(db),
		Rules:         NewRulerepositories(db),
		Stats:         NewStatsrepositories(db),
		Submissions:   NewSubmissionrepositories(db),
		Locations:     NewLocationRepository(db),
		Organizations: NewOrganizationRepository(db),
		Honey:         NewHoneyRepository(db),
	}
}
