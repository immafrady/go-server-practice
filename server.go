package main

import (
	"flag"
	"fradyspace.com/go-server-practice/app/service"
	"fradyspace.com/go-server-practice/router"
	_struct "fradyspace.com/go-server-practice/struct"
	"log"
	"net/http"
)

var (
	port    = flag.String("p", "3000", "设置端口号(默认3000)")
	isLocal = flag.Bool("l", false, "是否本地服务")
	env     = flag.String("env", "dev", "启动环境")
	address = "0.0.0.0"
)

func main() {
	if *isLocal {
		address = "127.0.0.1"
	}

	// 初始化数据库
	db := service.OpenDb(_struct.Env(*env))
	service.InitDb(db)
	router.Routers()
	log.Printf("Server Started at %s:%s!\n", address, *port)
	err := http.ListenAndServe(address+":"+*port, nil)
	if err != nil {
		log.Fatalln("Unable To Start: " + err.Error())
	}
}
