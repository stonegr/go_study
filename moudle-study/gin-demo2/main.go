package main

import (
	cs1 "cs"
	"fmt"
	controllers "gin-demos/controller"
)

func main() {
	// r := gin.Default()

	// r.GET("/", func(c *gin.Context) {
	// 	// c.String(http.StatusOK, "hello world")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"a": "b",
	// 	})
	// })
	// v := r.Group("/lt")
	// {
	// 	v.GET("/cs", controller.IndexController)
	// }

	// r.Run()
	fmt.Println("main")
	cs1.Cs()
	controllers.Cs()
}
