package main

import (
	"log"
	"my-web/controller/server_back"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetCurrentDirectory() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(filepath.Dir(os.Args[0])))
	if err != nil {
		log.Fatal(err)
	}

	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1)
}

func CreatRouter(g *gin.Engine) {
	// 服务器上传
	g.MaxMultipartMemory = 500 << 20 // 64 MiB
	server_route := g.Group("/server_back")
	server_route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok " + GetCurrentDirectory(),
		})
	})
	server_route.POST("/", server_back.UploadServerFile)
}
