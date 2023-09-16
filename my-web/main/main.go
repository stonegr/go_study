package main

// 导入gin包
import (
	"fmt"
	"io"
	"my-web/init_process"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
)

func init() {
	// 日志
	if _, err := os.Stat(init_process.Myconfig.LogDir); os.IsNotExist(err) {
		os.Mkdir(init_process.Myconfig.LogDir, 0777)
	}
	f := init_process.InitLogger(init_process.Myconfig.LogDir)
	defer f.Close()
	fmt.Println(init_process.DeferCs())
}

// 入口函数
func main() {
	// 初始化一个http服务对象
	f := &lumberjack.Logger{
		// 日志输出文件路径
		Filename: path.Join(init_process.Myconfig.LogDir, "output.log"),
		// 日志文件最大 size, 单位是 MB
		MaxSize: 1, // megabytes
		// 最大过期日志保留的个数
		MaxBackups: 3,
		// 保留过期文件的最大时间间隔,单位是天
		MaxAge: 1, //days
		// 是否需要压缩滚动日志, 使用的 gzip 压缩
		Compress: false, // disabled by default
	}
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()

	// r := gin.New()
	// r.Use(gin.Recovery())
	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	// 自定义日志输出格式
	// 	var statusColor, methodColor, resetColor string
	// 	if param.IsOutputColor() {
	// 		statusColor = param.StatusCodeColor()
	// 		methodColor = param.MethodColor()
	// 		resetColor = param.ResetColor()
	// 	}

	// 	if param.Latency > time.Minute {
	// 		param.Latency = param.Latency.Truncate(time.Second)
	// 	}
	// 	msg := fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
	// 		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
	// 		statusColor, param.StatusCode, resetColor,
	// 		param.Latency,
	// 		param.ClientIP,
	// 		methodColor, param.Method, resetColor,
	// 		param.Path,
	// 		param.ErrorMessage,
	// 	)
	// 	log.Infoln(msg)
	// 	return ""
	// }))

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
