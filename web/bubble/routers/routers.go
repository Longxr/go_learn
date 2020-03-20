package routers

import (
	"github.com/gin-gonic/gin"
	"web/bubble/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//加载静态文件
	r.Static("/static", "static")
	//加载模板文件
	r.LoadHTMLGlob("./templates/*")
	r.GET("/", controller.IndexHandler)

	//v1
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CreateATodo)
		//查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
