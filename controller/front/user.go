package front

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {

	fmt.Println("userService index")

	c.JSON(200, gin.H{
		"message": "hello userService index",
	})
}

func IndexPage(c *gin.Context) {

	fmt.Println("userService IndexPage")

	c.HTML(http.StatusOK, "frontIndex.html", gin.H{
		"title": "adminService website",
	})
}
