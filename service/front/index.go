package front

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {

	fmt.Println("front index")

	c.JSON(200, gin.H{
		"message": "hello front index",
	})
}

func IndexPage(c *gin.Context) {

	fmt.Println("front IndexPage")

	c.HTML(http.StatusOK, "frontIndex.html", gin.H{
		"title": "admin website",
	})
}
