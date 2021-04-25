package main

import (
	"ginBlog/config"
	"ginBlog/model"
	"ginBlog/routes"
)

func main() {
	//引用数据库
	model.InitDb()
	r := routes.InitRouter()
	r.Run(config.HttpPort)
}
