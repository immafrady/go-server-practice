package dao

import (
	"errors"
	"fradyspace.com/go-server-practice/models"
	"fradyspace.com/go-server-practice/utils/db"
)

func NewTag(name string) (err error) {
	gormDb := db.GetMysqlDB()

	tag := models.Tag{Name: name}
	if gormDb.NewRecord(tag) {
		gormDb.Create(&tag)
		if gormDb.NewRecord(tag) {
			err = errors.New("无法新增")
		}
	}

	return
}

func DeleteTag(name string) (err error) {
	gorm := db.GetMysqlDB()

	tag := models.Tag{Name: name}
	err = gorm.Unscoped().Delete(&tag).Error

	return
}

func GetTagId(name string) (id uint, err error) {
	gorm := db.GetMysqlDB()

	tag := models.Tag{Name: name}
	err = gorm.First(&tag).Error
	if err != nil {
		return
	}
	id = tag.ID
	return
}

func GetPostDetail() {

}

func UpdatePostDetail() {

}

func CreateNewPost() {

}

func DeletePost(id uint) {

}
