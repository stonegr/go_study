package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMysql() (err error) {
	dsn := "root:123-=shi@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return DB.Error
}

func main() {

	// 连接数据库
	err := initMysql()
	if err != nil {
		panic(err)
	}
	// 模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	// 标识模版文件引用依赖路径
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/fonts", "./static/fonts")
	r.Static("/static", "./static")
	// 标识模版文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	// 添加
	v1Group.POST("/todo", func(c *gin.Context) {
		var todo Todo
		c.BindJSON(&todo)
		err = DB.Create(&todo).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": "00000000000",
				"msg":  "Success",
				"date": todo,
			})
		}
	})

	// 查看所有
	v1Group.GET("/todo", func(c *gin.Context) {
		var todoList []Todo
		err := DB.Find(&todoList).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, todoList)
		}
	})

	// 查看某一个
	v1Group.GET("/todo/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
			return
		}
		var todo Todo
		err = DB.Where("id=?", id).First(&todo).Error
		if err == nil {
			c.JSON(http.StatusOK, todo)
		} else {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}
	})

	// 修改
	v1Group.PUT("/todo/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
			return
		}
		var todo Todo
		err = DB.Where("id=?", id).First(&todo).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.BindJSON(&todo)
		if err = DB.Save(&todo).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, todo)
		}
	})

	// 删除
	v1Group.DELETE("/todo/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
			return
		}
		err = DB.Where("id=?", id).Delete(&Todo{}).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{id: "deleted"})
		}

	})
	r.Run()
}
