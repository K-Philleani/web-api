package service

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"math/rand"
	"time"
	"web-api/dao"
	"web-api/model"
	"web-api/utils"
)

type SmsCodeService struct {

}


func (ss *SmsCodeService) SendCode(phone string) bool{
	// 1.产生一个验证码
	code := fmt.Sprintf("%06v",rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000))
	// 2.调用阿里云SDK, 完成发送
	config := utils.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AccessKey, config.AccessSecret)
	if err != nil {
		log.Println(err)
		return false
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	// 3.接收返回结果，判断发送状态
	response, err := client.SendSms(request)
	log.Println(response)
	if err != nil {
		return false
	}
	if response.Code == "OK" {
		sms := model.SmsCode{
			Phone:      phone,
			BizId:      response.BizId,
			Code:       code,
			CreateTime: time.Now().Unix(),
		}
		smsDao := new(dao.SmsCodeDao)
		row := smsDao.InsertCode(sms)
		return row>0
	}

	return false
}

