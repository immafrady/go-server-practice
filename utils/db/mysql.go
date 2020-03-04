package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	mysqlDB *gorm.DB
)

func StartMysql(dst string) (err error) {
	mysqlDB, err = gorm.Open("mysql", dst)
	return
}

func GetMysqlDB() (db *gorm.DB) {
	db = mysqlDB
	return
}
