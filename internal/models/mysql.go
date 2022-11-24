package models

import (
	"time"

	"github.com/solost23/tools/logger"
	"github.com/solost23/tools/mysql"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func NewMysqlConnect() (connect *gorm.DB) {
	var err error
	connect, err = mysql.NewMysqlConnect(&mysql.Config{
		UserName:        viper.GetString("connections.mysql.push_service.user"),
		Password:        viper.GetString("connections.mysql.push_service.password"),
		Host:            viper.GetString("connections.mysql.push_service.host"),
		Port:            viper.GetInt("connections.mysql.push_service.port"),
		DB:              viper.GetString("connections.mysql.push_service.db"),
		Charset:         viper.GetString("connections.mysql.push_service.charset"),
		MaxIdleConn:     10,
		MaxOpenConn:     100,
		ConnMaxLifeTime: time.Hour,
		Logger:          logger.NewMysqlLogger(),
	})
	if err != nil {
		panic(err)
	}
	return connect
}
