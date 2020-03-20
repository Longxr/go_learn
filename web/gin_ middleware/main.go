package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	name, ok := c.Get("name")	//在上下文中取值
	if !ok {
		name = "匿名用户"
	}
	fmt.Println("index")
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

//定义中间件,统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	//计时
	start := time.Now()
	c.Next() //调用后续处理函数
	//c.Abort() //阻止后续函数调用
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out ...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Set("name", "longxr")	//在上下文中设置值
	c.Next() //调用后续处理函数
	//c.Abort() //阻止后续函数调用
	fmt.Println("m2 out ...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			//if isLogin {
			//	c.Next()
			//} else {
			//	c.Abort()
			//}
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(m1, m2, authMiddleware(false))	//全局注册中间件函数
	r.GET("/login", indexHandler)

	r.Run(":9090")
}
