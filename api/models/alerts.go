package models

import "github.com/lib/pq"

type AlertLanguages struct {
	Page      string         `db:"page"`
	Languages pq.StringArray `db:"languages"`
}
