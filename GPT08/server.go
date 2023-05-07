package GPT08

import (
	"fmt"
	"golang-practise/GPT08/config"
	"golang-practise/GPT08/model"
	"net/http"
	"strconv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			// 可以在这里记录错误日志
			c.JSON(http.StatusInternalServerError, model.NewError(500, "服务器出错了"))
		}
	}
}

func ServerRun() {
	config.LoadConfig()
	config.InitDB()

	r := gin.Default()
	r.Use(ErrorHandler())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("hello", hello)
	r.GET("todos", todolist)
	r.GET("todo", todo)
	r.Run(":" + viper.GetString("server.port"))
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

func todolist(c *gin.Context) {
	var todos []model.Todo
	config.GlobalDb.Order("due_timestamp DESC").Find(&todos)
	c.JSON(200, todos)
}

func todo(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.Error(fmt.Errorf("查询失败"))
	}
	var todo model.Todo
	config.GlobalDb.First(&todo, id)
	c.JSON(200, todo)
}
