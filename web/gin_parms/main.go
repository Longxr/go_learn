package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pwd"`
}

func main() {
	r := gin.Default()

	//获取URI路径参数
	r.GET("uri/:name/:age", func(c *gin.Context){
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age": age,
		})
	})

	r.GET("/user", func(c *gin.Context){
		var u UserInfo
		err := c.ShouldBind(&u)	//通过反射将参数字段映射到结构体
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "get ok",
			})
		}

	})

	r.POST("/user", func(c *gin.Context){
		var u UserInfo
		err := c.ShouldBind(&u)	//通过反射将参数字段映射到结构体
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "post ok",
			})
		}

	})

	r.Run(":9090")
}
