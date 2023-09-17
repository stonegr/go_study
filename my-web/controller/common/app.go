package common

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetXrayLl(c *gin.Context) {
	// output, code := ExecShell(init_process.XRAYSCRIPTPATH)
	output, code := ExecShell("ls")
	if code == 0 {
		c.String(http.StatusOK, fmt.Sprintf("已发送,output:%s\n", output))
	} else {
		c.String(http.StatusOK, fmt.Sprintf("发送失败,原因%s\n", output))
	}
}
