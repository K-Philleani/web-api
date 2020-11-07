package main

import (
	"github.com/gin-gonic/gin"
	"web-api/Controller"
	"web-api/utils"
)

func main() {
	acf, err := utils.ParseConfig("./config/app.json")
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	registerRouter(router)
	router.Run(acf.AppHost + ":" + acf.AppPort)
}

// 路由配置
func registerRouter(router *gin.Engine) {
	new(Controller.SmsCodeController).Router(router)
}
