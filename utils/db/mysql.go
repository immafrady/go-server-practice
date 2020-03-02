package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlDB *sql.DB
)

func StartMysql(dst string) (err error) {
	mysqlDB, err = sql.Open("mysql", dst)
	return
}

func GetMysqlDB() (db *sql.DB) {
	db = mysqlDB
	return
}
