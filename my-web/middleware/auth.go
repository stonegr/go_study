package middleware

import (
	"fmt"
	"my-web/init_process"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerBackAuth0(c *gin.Context) {
	fmt.Println("0 start")
}

func ServerBackAuth(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	if pwd == "" || init_process.Myconfig.MachineAuth[name] != pwd {
		c.String(http.StatusForbidden, "未授权")
		c.Abort()
	}
}

func ServerBackAuth2(c *gin.Context) {
	fmt.Println("2 start")
}
