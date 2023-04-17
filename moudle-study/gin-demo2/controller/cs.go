package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexController(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}

func Cs() {
	fmt.Println("controller-controllers-cs")
}
