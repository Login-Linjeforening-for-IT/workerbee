package models

type PathLanguages struct {
	ID        int    `db:"id" json:"id"`
	Page      string `db:"page" json:"page"`
	Languages string `db:"language" json:"language"`
}

type PathLanguagesWithCount struct {
	PathLanguages
	TotalCount int `db:"total_count" json:"-"`
}

type HoneyContent struct {
	Language string `db:"language"`
	Text     string `db:"text"`
}

type CreateHoney struct {
	ID       int    `db:"id" json:"id"`
	Service  string `db:"service" json:"service" validate:"required"`
	Page     string `db:"page" json:"page" validate:"required"`
	Language string `db:"language" json:"language" validate:"required,oneof=en no"`
	Text     string `db:"text" json:"text" validate:"required"`
}

type LanguageContentResponse struct {
	Service  string                       `db:"service" json:"service"`
	Page     string                       `db:"page" json:"page"`
	Language string                       `db:"language" json:"language"`
	Text     map[string]map[string]string `db:"text" json:"text"`
}
