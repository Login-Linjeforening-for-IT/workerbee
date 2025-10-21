package models

import "github.com/lib/pq"

type PathLanguages struct {
	Page      string         `db:"page"`
	Languages pq.StringArray `db:"languages"`
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
