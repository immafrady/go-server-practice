package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
)

var (
	Config *config
	Pwd    string
	env    string
)

type config struct {
	Username string `json:"mysql_username"`
	Password string `json:"mysql_password"`
	Database string `json:"mysql_database"`
	Host     string `json:"mysql_host"`
	Port     string `json:"mysql_port"`
}

func init() {
	flag.StringVar(&env, "env", "dev", "环境")
	Pwd, _ = os.Getwd()
}

func InitConfig() (err error) {
	str, err := ioutil.ReadFile(Pwd + "/config/" + env + ".json")
	if err != nil {
		return
	}

	Config = &config{}

	err = json.Unmarshal(str, Config)
	if err != nil {
		return
	}
	return
}
