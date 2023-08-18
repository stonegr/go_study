package server_back

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func UploadServerFile(c *gin.Context) {
	// 获取上传文件，返回的是multipart.FileHeader对象，代表一个文件，里面包含了文件名之类的详细信息
	// file是表单字段名字
	dir_name := c.PostForm("name")
	file, _ := c.FormFile("file")
	// 打印上传的文件名
	log.Println(dir_name, file.Filename)

	// 将上传的文件，保存到./data/1111.jpg 文件中
	err := c.SaveUploadedFile(file, fmt.Sprintf("./server_back/%s/%s", dir_name, file.Filename))
	if err != nil {
		panic(err)
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
