package hero

import (
	"fradyspace.com/go-server-practice/app/service"
	"fradyspace.com/go-server-practice/utils"
)

var logger *utils.Logger

func init() {
	logger = utils.NewLogger("hero service")
}

func QueryHeroList() (data string, err error)  {
	sql := `SELECT id, name FROM heroes`
	db := service.GetCurrentDb()
	r, err := db.Query(sql)
	if err != nil {
		return
	}
	defer r.Close()
	for r.Next() {
		r.
	}
}