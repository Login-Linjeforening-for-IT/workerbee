package repositories

import (
	"encoding/json"
	"os"
	"workerbee/db"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
)

type Formrepositories interface {
	GetForms(search, limit, offset, orderBy, sort string) ([]models.FormWithTotalCount, error)
	GetForm(id string) (*models.FormWithQuestion, error)
	PostForm(form models.Form) (models.Form, error)
	PutForm(id string, form models.Form) (models.Form, error)
	DeleteForm(id string) (models.Form, error)
}

type formrepositories struct {
	db *sqlx.DB
}

func NewFormrepositories(db *sqlx.DB) Formrepositories {
	return &formrepositories{db: db}
}

func (r *formrepositories) GetForms(search, limit, offset, orderBy, sort string) ([]models.FormWithTotalCount, error) {
	forms, err := db.FetchAllElements[models.FormWithTotalCount](
		r.db,
		"./db/forms/get_forms.sql",
		orderBy, sort,
		limit,
		offset,
		search,
	)
	if err != nil {
		return nil, err
	}

	return forms, nil
}

func (r *formrepositories) GetForm(id string) (*models.FormWithQuestion, error) {
	type formWithQuestionsRaw struct {
		models.FormWithQuestion
		QuestionsRaw json.RawMessage `db:"questions"`
		UserRaw      json.RawMessage `db:"user"`
	}
	f := formWithQuestionsRaw{}

	sqlBytes, err := os.ReadFile("./db/forms/get_form.sql")
	if err != nil {
		return nil, err
	}

	query := string(sqlBytes)
	err = r.db.Get(&f, query, id)
	if err != nil {
		return nil, err
	}

	if len(f.QuestionsRaw) > 0 {
		_ = json.Unmarshal(f.QuestionsRaw, &f.Questions)
	}
	if len(f.UserRaw) > 0 {
		var user models.User
		_ = json.Unmarshal(f.UserRaw, &user)
		f.User = &user
	}

	return &f.FormWithQuestion, nil
}

func (r *formrepositories) PostForm(form models.Form) (models.Form, error) {
	newForm := models.Form{}

	sqlBytes, err := os.ReadFile("./db/forms/post_form.sql")
	if err != nil {
		return models.Form{}, err
	}

	query := string(sqlBytes)
	err = r.db.Get(&newForm, query, form.UserID, form.Title, form.Description, form.Capacity, form.OpenAt, form.CloseAt)
	if err != nil {
		return models.Form{}, err
	}

	return newForm, nil
}

func (r *formrepositories) PutForm(id string, form models.Form) (models.Form, error) {
	updatedForm := models.Form{}

	sqlBytes, err := os.ReadFile("./db/forms/put_form.sql")
	if err != nil {
		return models.Form{}, err
	}
	query := string(sqlBytes)
	err = r.db.Get(&updatedForm, query, id, form.Title, form.Description, form.Capacity, form.OpenAt, form.CloseAt)
	if err != nil {
		return models.Form{}, err
	}

	return updatedForm, nil
}

func (r *formrepositories) DeleteForm(id string) (models.Form, error) {
	deletedForm := models.Form{}

	sqlBytes, err := os.ReadFile("./db/forms/delete_form.sql")
	if err != nil {
		return models.Form{}, err
	}
	query := string(sqlBytes)

	err = r.db.Get(&deletedForm, query, id)
	if err != nil {
		return models.Form{}, err
	}

	return deletedForm, nil
}
