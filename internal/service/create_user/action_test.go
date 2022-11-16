package create_user

import (
	"context"
	"fmt"
	"github.com/solost23/go_interface/gen_go/common"
	"github.com/solost23/go_interface/gen_go/user_service"
	"github.com/solost23/tools/logger"
	"github.com/solost23/tools/mysql"
	"github.com/spf13/viper"
	"testing"
	"time"
)

var (
	WebConfigPath = "../../../configs/conf.yml"
)

func InitConfig() {
	viper.SetConfigFile(WebConfigPath)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func TestAction_Deal(t *testing.T) {
	InitConfig()
	connect, err := mysql.NewMysqlConnect(
		&mysql.Config{
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
	client := NewActionWithCtx(context.Background())
	client.SetHeader(&common.RequestHeader{})
	client.SetMysql(connect)
	// 一个run 就是一个case
	t.Run("测试首次创建", func(t *testing.T) {
		request := &user_service.CreateUserRequest{
			Header: &common.RequestHeader{},
			User: &user_service.User{
				UserName: "ty1",
				Password: "123",
			},
		}
		resp, err := client.Deal(context.Background(), request)
		if err != nil {
			t.Error(err)
			return
		}
		if resp.ErrorInfo.Code != 0 {
			t.Error("resp.ErrorInfo.Code:", resp.ErrorInfo.Msg)
			return
		}
		// 核对数据正确性, 从数据库中查找数据，如果和预期相符

		// 数据正常将数据库中数据清理掉

	})
	t.Run("测试库内已经存在已删除数据", func(t *testing.T) {
		fmt.Println("测试库内已经存在已删除数据")
	})
	t.Run("测试库内已存在数据", func(t *testing.T) {
		fmt.Println("测试库内已存在数据")
	})
}
