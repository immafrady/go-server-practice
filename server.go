package main

import (
	"flag"
	. "fradyspace.com/go-server-practice/config"
	"fradyspace.com/go-server-practice/controllers"
	"fradyspace.com/go-server-practice/models"
	"fradyspace.com/go-server-practice/utils/db"
	"fradyspace.com/go-server-practice/utils/logger"
	"log"
	"net/http"
)

var (
	port = flag.String("p", "3000", "启动端口号")
)

func main() {
	flag.Parse()
	log.Println("初始化配置文件...")
	logger.CheckError("初始化配置失败", InitConfig())
	log.Println("初始化配置文件[完成]")

	log.Println("初始化数据库...")
	logger.CheckError("连接数据库失败", db.StartMysql(Config.Username+":"+Config.Password+"@tcp("+Config.Host+":"+Config.Port+")/"+Config.Database))
	logger.CheckError("初始化数据库失败", models.InitDatabase())
	log.Println("初始化数据库[完成]")

	log.Printf("启动服务[端口: %s]...\n", *port)
	controllers.StartRouting()
	logger.CheckError("启动服务失败", http.ListenAndServe(":"+*port, nil))
}
