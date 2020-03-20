package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//外部重定向，浏览器地址改变
	r.GET("/redit", func(c *gin.Context){
		c.Redirect(http.StatusMovedPermanently,"https://www.baidu.com")
	})

	//内部重定向，浏览器地址不变
	r.GET("/a", func(c *gin.Context){
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})

	r.GET("/b", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "bbbbb",
		})
	})
	
	//Any 请求方法集合
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
		}
	})

	//404
	r.NoRoute(func(c *gin.Context) {
		//c.JSON(http.StatusNotFound, gin.H{"msg": "longxuan.ren"})
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<p> 2333333</p>`)
	})

	//路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
		})
		videoGroup.GET("xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
		})
	}

	r.Run(":9090")
}
