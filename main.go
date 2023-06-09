package main

import (
	"giligili/app"
	"giligili/app/http/server"
	"log"
)

func main() {
	err := app.Init() //加载静态配置文件
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(app.Config)
	err = server.NewServer()
	if err != nil {
		log.Println(err)
	}
}
