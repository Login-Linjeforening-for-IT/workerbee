package repositories

import (
	"github.com/jmoiron/sqlx"
	"workerbee/models"
)

type Repositories struct {
	Forms       Formrepositories
	Events      Eventrepositories
	Jobs        Jobsrepositories
	Questions   Questionrepositories
	Rules       Rulerepositories
	Stats       Statsrepositories
	Submissions Submissionrepositories
}

// DeleteEvent implements Eventrepositories.
func (r *Repositories) DeleteEvent(id int) (models.Event, error) {
	panic("unimplemented")
}

// GetEvent implements Eventrepositories.
func (r *Repositories) GetEvent(id int) (models.Event, error) {
	panic("unimplemented")
}

// GetEvents implements Eventrepositories.
func (r *Repositories) GetEvents(search string, limit string, offset string, orderBy string, sort string, historical bool) ([]models.EventWithTotalCount, error) {
	panic("unimplemented")
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
	}
}
