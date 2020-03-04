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

func ExecDb(query string, args ...interface{}) (result sql.Result, err error) {
	db := GetMysqlDB()
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	result, err = stmt.Exec(args...)
	return
}
