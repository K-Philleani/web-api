package main

import (
	"github.com/gin-gonic/gin"
	"web-api/utils"
)

func main() {
	acf, err := utils.ParseConfig("./config/app.json")
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"code": 1,
		})
	})
	router.Run(acf.AppHost + ":" + acf.AppPort)
}
