package GPT08

import (
	"golang-practise/GPT08/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ServerRun() {
	config.LoadConfig()
	config.InitDB()

	r := gin.Default()
	r.GET("hello", hello)
	r.Run(":" + viper.GetString("server.port"))
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
