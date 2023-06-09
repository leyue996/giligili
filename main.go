package main

import (
	"giligili/app"
	"giligili/app/http/server"
	"log"
)

func main() {
	app.Init() //加载静态配置文件

	//fmt.Println(app.Config)
	err := server.NewServer()
	if err != nil {
		log.Println(err)
	}
}
