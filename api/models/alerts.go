package models

type Alert struct {
	ID            string `db:"id" json:"id"`
	Service       string `db:"service" json:"service"`
	Page          string `db:"page" json:"page"`
	TitleEn       string `db:"title_en" json:"title_en"`
	TitleNo       string `db:"title_no" json:"title_no"`
	Description   string `db:"description_en" json:"description_en"`
	DescriptionNo string `db:"description_no" json:"description_no"`
}
