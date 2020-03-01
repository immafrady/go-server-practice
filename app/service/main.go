package service

import (
	"database/sql"
	_struct "fradyspace.com/go-server-practice/struct"
	"fradyspace.com/go-server-practice/utils"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

var logger *utils.Logger
var db *sql.DB

func init() {
	logger = utils.NewLogger("base sql")
}

func OpenDb(env _struct.Env) *sql.DB {
	if db != nil {
		err := db.Ping()
		if err != nil {
			goto INIT
		}
		return db
	}

INIT:
	config := utils.LoadConfig(env).MySQL

	db, err := sql.Open("mysql", config.Username+":"+config.Password+"@tcp("+config.Host+":"+config.Port+")/"+config.Database)
	if err != nil {
		logger.Fatalf(err.Error())
	}
	return db
}

func InitDb(db *sql.DB) {
	config, err := utils.GetPathConfig(_struct.PathInitSQL)
	if err != nil {
		logger.Errorf(err.Error())
	}

	str, err := ioutil.ReadFile(config)
	if err != nil {
		logger.Fatalf(err.Error())
	}
	_, err = db.Exec(string(str))
	if err != nil {
		logger.Fatalf(err.Error())
	}
}

func GetCurrentDb() *sql.DB {
	if db == nil {
		logger.Fatalf("db not init")
	}
	return db
}
