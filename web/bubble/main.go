package main

import (
	"web/bubble/dao"
	"web/bubble/models"
	"web/bubble/routers"
)

func main() {
	//创建数据库
	//sql: CREATE DATABASE bubble
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run(":9090")
}
