package main

import (
	"log"
	"my-web/controller/server_back"
	"my-web/middleware"
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
	g.MaxMultipartMemory = 500 << 20 // 500 MiB

	api_group := g.Group("/go-api")

	{
		server_back_route := api_group.Group("/server_back")
		{
			server_back_route.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"msg": "ok " + GetCurrentDirectory(),
				})
			})
			server_back_route.POST("/",
				// middleware.ServerBackAuth0,
				middleware.ServerBackAuth,
				// middleware.ServerBackAuth2,
				server_back.UploadServerFile)
		}
	}
}
