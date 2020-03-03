package models

type HeroModel struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CreateDate string `json:"create_date"`
	UpdateDate string `json:"update_date"`
}
