package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type BlogModel struct {
	gorm.Model
	Id        int            `json:"id"`
	Title     string         `json:"title"`
	Introduce string         `json:"introduce"`
	ImgUrl    sql.NullString `json:"img_url"`
}

func (BlogModel) TableName() string {
	return "blog_"
}

type TagModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
