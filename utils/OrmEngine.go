package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

)

type Orm struct {
	*gorm.DB
}

var DbOrm *gorm.DB


func OrmEngine(acf *AppConfig) error {
	var err error
	database := acf.DataBase
	DbOrm, err = gorm.Open(mysql.Open(database.Address), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	return nil
}