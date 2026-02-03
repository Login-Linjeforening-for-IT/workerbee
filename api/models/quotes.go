package models

import "workerbee/internal"

type BaseQuote struct {
	ID        int                `db:"id" json:"id"`
	Author    string             `db:"author" json:"-"`
	Quoted    string             `db:"quoted" json:"quoted"`
	Content   string             `db:"content" json:"content"`
	CreatedAt internal.LocalTime `db:"created_at" json:"created_at"`
	UpdatedAt internal.LocalTime `db:"updated_at" json:"updated_at"`
}

type QuoteWithTotalCount struct {
	BaseQuote
	TotalCount int `db:"total_count" json:"-"`
}
