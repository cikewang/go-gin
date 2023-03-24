package main

import (
	"github.com/jinzhu/configor"
	"go-gin/library"
	"go-gin/models"
	"go-gin/routers"
)

// 初始化项目
func init() {
	// 加载配置文件
	configor.Load(&library.Config, "./config/config.yaml")
	// 初始化 MySQL
	models.MySQLConnect()
}

func main() {
	r := routers.Router()

	r.Run(":8080")
}
