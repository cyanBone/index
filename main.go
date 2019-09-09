package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "你好师姐",
		})
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}
