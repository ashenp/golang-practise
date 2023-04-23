package config

import (
	"fmt"
	"golang-practise/GPT08/model"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var (
	db *gorm.DB
)

func InitDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.address"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
	)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&model.Todo{})
}
