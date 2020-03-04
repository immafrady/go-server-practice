package models

import "fradyspace.com/go-server-practice/utils/db"

func InitDatabase() (err error) {
	gormDb := db.GetMysqlDB()
	gormDb.Set("gorm:table_options", "").CreateTable(&BlogModel{}, &TagModel)
	return
}
