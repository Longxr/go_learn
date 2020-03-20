package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")

	//方法一：使用map
	r.GET("/json", func(c *gin.Context) {
		//data := map[string]interface{}{
		//	"name": "longxr",
		//	"message": "hello world",
		//	"age": 23,
		//}
		data := gin.H{
			"name": "longxr",
			"message": "hello world",
			"age": 23,
		}
		c.JSON(http.StatusOK, data)
	})

	//方法二：使用结构体
	type msg struct {
		Name string `json:"name"`
		Age int `json:"age"`
		Message string `json:"message"`
	}
	r.GET("json2", func(c *gin.Context){
		data := msg {
			Name: "longxr",
			Message: "hello world",
			Age: 23,
		}
		c.JSON(http.StatusOK, data)
	})

	//获取get的query string 参数
	r.GET("/query", func(c *gin.Context){
		//name := c.Query("user")
		name := c.DefaultQuery("user", "nobody")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	r.GET("/login", func(c *gin.Context){
		c.HTML(http.StatusOK, "login.html", nil)
	})
	//获取post的from参数
	r.POST("/login", func(c *gin.Context){
		//username := c.PostForm("username")
		//password := c.PostForm("password")
		username := c.DefaultPostForm("username", "somebody")
		password := c.DefaultPostForm("xxx", "*****")	//不填是空字符串，没有对应字段才用default
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name": username,
			"Password": password,
		})
	})

	r.Run(":9090")
}
