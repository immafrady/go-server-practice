package _struct

type Env string
type Path string

const (
	DEV  Env = "dev"
	PROD Env = "prod"

	PathInitSQL Path = "init_sql"
)

type Config struct {
	MySQL MySQL `json:"my_sql"`
}

type MySQL struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}
