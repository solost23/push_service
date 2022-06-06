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
		UserName:        viper.GetString("connections.mysql.hello_world.user"),
		Password:        viper.GetString("connections.mysql.hello_world.password"),
		Host:            viper.GetString("connections.mysql.hello_world.host"),
		Port:            viper.GetInt("connections.mysql.hello_world.port"),
		DB:              viper.GetString("connections.mysql.hello_world.db"),
		Charset:         viper.GetString("connections.mysql.hello_world.charset"),
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
