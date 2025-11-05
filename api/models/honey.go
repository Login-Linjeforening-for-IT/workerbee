package models

type PathLanguages struct {
	ID        int    `db:"id" json:"id"`
	Page      string `db:"page" json:"page"`
	Languages string `db:"language" json:"language"`
}

type PathLanguagesWithCount struct {
	PathLanguages
	TotalCount int `db:"total_count" json:"total_count"`
}

type HoneyContent struct {
	Language string `db:"language"`
	Text     string `db:"text"`
}

type LanguageContent struct {
	Service  string `db:"service" json:"service"`
	Page     string `db:"page" json:"page"`
	Language string `db:"language" json:"language"`
	Text     string `db:"text" json:"text"`
}

type LanguageContentResponse struct {
	Service  string                       `db:"service" json:"service"`
	Page     string                       `db:"page" json:"page"`
	Language string                       `db:"language" json:"language"`
	Text     map[string]map[string]string `db:"text" json:"text"`
}
