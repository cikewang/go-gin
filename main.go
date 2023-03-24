package main

import (
	"github.com/jinzhu/configor"
	"go-gin/library"
)

// 初始化项目
func init() {
	// 加载配置文件
	configor.Load(&library.Config, "./config/config.yaml")
}

func main() {
	r := library.Router()

	// fmt.Printf("config: %d", library.Config.Redis.Port)

	r.Run(":8080")
}
