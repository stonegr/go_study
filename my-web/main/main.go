package main

// 导入gin包
import (
	"fmt"
	"my-web/init_process"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	f := init_process.InitLogger()
	defer f.Close()
	log.Info("日志处理器初始化完成!")
	fmt.Println(init_process.DeferCs())
}

// 入口函数
func main() {
	// 初始化一个http服务对象
	r := gin.Default()

	// 设置一个get请求的路由，url为/ping, 处理函数（或者叫控制器函数）是一个闭包函数。
	r.GET("/", func(c *gin.Context) {
		// 通过请求上下文对象Context, 直接往客户端返回一个json
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	CreatRouter(r)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
