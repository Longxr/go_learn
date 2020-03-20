package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")

	//返回index页面
	r.GET("index", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("upload", func(c *gin.Context){
		//从请求中读取文件
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//dst := fmt.Sprintf("./%s", f.Filename)
			dst := path.Join("./", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
		//文件保存

	})

	r.Run(":9090")
}
