package server_back

import (
	"fmt"
	"my-web/init_process"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func UploadServerFile(c *gin.Context) {
	var server_up_fieled SERVER_UP_FIELED
	if err := c.ShouldBind(&server_up_fieled); err != nil {
		c.String(http.StatusForbidden, fmt.Sprint(err))
		log.Error(err)
		return
	}
	// log.Println(server_up_fieled.Name, server_up_fieled.File.Filename)

	file_path := fmt.Sprintf("./%s/%s/%s", init_process.Myconfig.BackDir, server_up_fieled.Name, server_up_fieled.File.Filename)

	// 验证文件是否存在
	if exsits, _ := PathExists(file_path); exsits {
		c.String(http.StatusForbidden, fmt.Sprintf("文件: %s, 已存在!", server_up_fieled.File.Filename))
		return
	}

	// 将上传的文件，保存到./data/1111.jpg 文件中
	if err := c.SaveUploadedFile(server_up_fieled.File, file_path); err != nil {
		c.String(http.StatusForbidden, fmt.Sprint(err))
		log.Error(err)
		return
	} else {
		log.Info(fmt.Sprintf("%s成功上传: %s", server_up_fieled.Name, server_up_fieled.File.Filename))
	}

	_ = DealWithFiles(path.Join(init_process.Myconfig.BackDir, server_up_fieled.Name), init_process.Myconfig.MaxFile)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", server_up_fieled.File.Filename))
}
