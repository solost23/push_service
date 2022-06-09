package test

import (
	"fmt"
	"github.com/solost23/tools/logger"
	"github.com/solost23/tools/mysql"
	"github.com/spf13/viper"
	"testing"
	"time"
)

var (
	WebConfigPath = "../configs/conf.yml"
)

func InitConfig() {
	viper.SetConfigFile(WebConfigPath)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func TestBatchGet(t *testing.T) {
	InitConfig()
	starTime := time.Now().Unix()
	t.Log("start:", starTime)
	conn, err := mysql.NewMysqlConnect(&mysql.Config{
		UserName:        viper.GetString("connections.mysql.my_grpc_frame.user"),
		Password:        viper.GetString("connections.mysql.my_grpc_frame.password"),
		Host:            viper.GetString("connections.mysql.my_grpc_frame.host"),
		Port:            viper.GetInt("connections.mysql.my_grpc_frame.port"),
		DB:              viper.GetString("connections.mysql.my_grpc_frame.db"),
		Charset:         viper.GetString("connections.mysql.my_grpc_frame.charset"),
		MaxIdleConn:     10,
		MaxOpenConn:     100,
		ConnMaxLifeTime: time.Hour,
		Logger:          logger.NewMysqlLogger(),
	})
	if err != nil {
		panic(err)
	}
	tx := conn.Begin()
	var datas []map[string]interface{}
	for i := 0; i <= 100000; i++ {
		fmt.Printf("查找第%d次 \n", i)
		tx.Raw("SELECT * FROM user").Find(&datas)
	}
	tx.Commit()
	t.Log(datas)
	t.Log(err)
}
