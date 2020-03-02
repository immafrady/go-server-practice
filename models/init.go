package models

import (
	"fradyspace.com/go-server-practice/config"
	"fradyspace.com/go-server-practice/utils/db"
	"io/ioutil"
)

func InitDatabase() (err error) {
	str, err := ioutil.ReadFile(config.Pwd + "/sql/init.sql")
	if err != nil {
		return
	}
	database := db.GetMysqlDB()
	_, err = database.Exec(string(str))
	return
}
