package repositories

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Forms         Formrepositories
	Categories    CategoireRepository
	Locations     LocationRepository
	Organizations OrganizationRepository
	Events        Eventrepositories
	Jobs          Jobsrepositories
	Questions     Questionrepositories
	Rules         Rulerepositories
	Stats         Statsrepositories
	Submissions   Submissionrepositories
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Forms:       NewFormrepositories(db),
		Events:      NewEventrepositories(db),
		Jobs:        NewJobrepositories(db),
		Questions:   NewQuestionrepositories(db),
		Rules:       NewRulerepositories(db),
		Stats:       NewStatsrepositories(db),
		Submissions: NewSubmissionrepositories(db),
		Categories: NewCategoireRepository(db),
		Locations: NewLocationRepository(db),
		Organizations: NewOrganizationRepository(db),
	}
}
