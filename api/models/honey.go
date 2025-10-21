package models

import "github.com/lib/pq"

type PathLanguages struct {
	Page      string         `db:"page" json:"page"`
	Languages pq.StringArray `db:"languages" json:"languages"`
}
