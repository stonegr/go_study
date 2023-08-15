package main

import "github.com/gin-gonic/gin"

func CreatRouter(g *gin.Engine) {
	server_route := g.Group("/server_back")
	server_route.POST("/")
}
