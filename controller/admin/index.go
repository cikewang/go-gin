package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/models"
	"net/http"
)

func Index(c *gin.Context) {

	fmt.Println("adminService index")

	c.HTML(http.StatusOK, "adminIndex.html", gin.H{
		"title": "adminService website",
	})

}

func GetAdminInfo(c *gin.Context) {
	adminId := 1
	admin, err := models.GetAdminByAdminID(adminId)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    10001,
			"message": "管理员不存在",
		})
	} else {
		c.JSON(200, gin.H{
			"code":    10000,
			"message": "OK",
			"data":    admin,
		})
	}

}
