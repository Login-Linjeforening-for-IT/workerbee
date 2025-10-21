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
