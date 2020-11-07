package Controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"web-api/service"
)

type SmsCodeController struct {

}

func (sc *SmsCodeController) Router(router *gin.Engine) {
	router.GET("sendCode", sc.SendSmsCode)
}

func (sc *SmsCodeController) SendSmsCode(c *gin.Context) {
	phone, exist := c.GetQuery("phone")
	if !exist {
		c.JSON(200, map[string]interface{}{
			"code": 0,
			"message": "参数解析失败",
		})
	}
	log.Println(phone)
	service := new(service.SmsCodeService)
	isSend := service.SendCode(phone)
	if isSend {
		c.JSON(200, map[string]interface{}{
			"code": 1,
			"message": "发送成功",
		})
	} else {
		c.JSON(200, map[string]interface{}{
			"code": 0,
			"message": "发送失败",
		})
	}
}