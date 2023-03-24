package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin/controller/admin"
	"go-gin/controller/front"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("./templates/**/*")
	r.Static("/static", "./static")

	// 前端路由
	frontGroup := r.Group("/front")
	{
		// json API
		frontGroup.GET("/index", front.Index)
		// HTML
		frontGroup.GET("/page", front.IndexPage)
	}

	// 后台管理路由
	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/page", admin.Index)
		adminGroup.GET("/admin", admin.GetAdminInfo)
	}

	return r
}
