package dao

import (
	"web-api/model"
	"web-api/utils"
)

type SmsCodeDao struct {

}

func (scd *SmsCodeDao) InsertCode(sms model.SmsCode) int64{
	res := utils.DbOrm.Create(&sms)
	return res.RowsAffected
}
