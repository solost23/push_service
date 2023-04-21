package send_email

import (
	"context"
	"fmt"
	"github.com/solost23/protopb/gen/go/protos/push"
	"github.com/spf13/viper"
	"push_service/configs"
	"push_service/internal/models"
	"testing"
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
	mdb, _ := models.InitMysql(&configs.MySQLConf{})
	type arg struct {
		topic       string
		name        string
		addr        string
		contentType string
		content     string
	}
	type want struct {
		err error
	}
	type test struct {
		name string
		ctx  context.Context
		arg  arg
		want want
	}
	tests := []test{
		{
			name: "case 1",
			ctx:  context.Background(),
			arg: arg{
				topic:       "测试",
				name:        "",
				addr:        "280***@qq.com",
				contentType: "text/plain",
				content:     "测试发送邮件",
			},
			want: want{
				err: nil,
			},
		},
		//{},
		//{},
	}
	client := NewActionWithCtx(context.Background())
	client.SetMysql(mdb)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := &push.SendEmailRequest{
				Email: &push.Email{
					Topic:       test.arg.topic,
					Name:        test.arg.name,
					Addr:        test.arg.addr,
					ContentType: test.arg.contentType,
					Content:     test.arg.content,
				},
			}
			_, err := client.Deal(test.ctx, request)
			if err != test.want.err {
				t.Error(err)
				return
			}
			fmt.Println("发送成功")
		})
	}
}
