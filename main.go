package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"code": 1,
		})
	})
	router.Run(":9091")
}
