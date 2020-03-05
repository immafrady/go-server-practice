package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title     string         `gorm:"not null"`
	Introduce string         `gorm:"type:mediumtext"`
	ImgUrl    sql.NullString `gorm:"type:tinytext"`
	Tags      []Tag          `gorm:"many2many:posts_tags;association_foreignkey"`
}

func (Post) TableName() string {
	return "blog_posts"
}

type Tag struct {
	gorm.Model
	Name  string `gorm:"UNIQUE"`
	Posts []Post `gorm:"many2many:posts_tags"`
}

func (Tag) TableName() string {
	return "blog_tags"
}
