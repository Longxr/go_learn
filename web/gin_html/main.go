package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	//加载静态文件
	r.Static("/static", "./static")//第一个路径是模板中写的相对路径名，第二个路径是文件实际相对路径
	//gin框架中添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("templates/**/*")	//模板解析,**代表目录，*代表文件
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{	//模板渲染
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{	//模板渲染
			"title": "users/index",
			"link": "<a href='https://longxuan.ren'>longxr的博客</a>",
		})
	})

	r.GET("users/4449", func(c *gin.Context) {
		c.HTML(http.StatusOK, "moban4449/index.html", nil)
	})

	r.Run(":9090")
}
