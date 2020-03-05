package models

import "fradyspace.com/go-server-practice/utils/db"

func InitDatabase() (err error) {
	gormDb := db.GetMysqlDB()
	gormDb = gormDb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT charset = utf8mb4")
	dbs := []interface{}{
		&Tag{},
		&Post{},
	}
	for _, model := range dbs {
		if !gormDb.HasTable(model) {
			gormDb.CreateTable(model)
		} else {
			gormDb.AutoMigrate(model)
		}
	}
	return
}
