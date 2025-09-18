package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
)

type FormRepository interface {
	GetForms(search, limit, offset, orderBy, sort string) ([]models.FormWithTotalCount, error)
	GetForm(id string) (*models.FormWithQuestion, error)
	PostForm(form models.Form) (models.Form, error)
	PatchForm(id string, form models.Form) (models.Form, error)
	DeleteForm(id string) (models.Form, error)
}

type formRepository struct {
	db *sqlx.DB
}

func NewFormRepository(db *sqlx.DB) FormRepository {
	return &formRepository{db: db}
}

func (r *formRepository) GetForms(search, limit, offset, orderBy, sort string) ([]models.FormWithTotalCount, error) {
	forms := []models.FormWithTotalCount{}

	sqlBytes, err := os.ReadFile("./db/forms/get_forms.sql")
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("%s ORDER BY %s %s\nLIMIT $2 OFFSET $3;", string(sqlBytes), orderBy, sort)

	err = r.db.Select(&forms, query, search, limit, offset)
	if err != nil {
		return nil, err
	}

	return forms, nil
}

func (r *formRepository) GetForm(id string) (*models.FormWithQuestion, error) {
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

func (r *formRepository) PostForm(form models.Form) (models.Form, error) {
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

func (r *formRepository) PatchForm(id string, form models.Form) (models.Form, error) {
	updatedForm := models.Form{}

	sqlBytes, err := os.ReadFile("./db/forms/patch_form.sql")
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

func (r *formRepository) DeleteForm(id string) (models.Form, error) {
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