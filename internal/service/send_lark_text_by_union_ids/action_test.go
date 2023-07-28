package send_lark_text_by_union_ids

import (
	"context"
	"fmt"
	"testing"

	"github.com/solost23/protopb/gen/go/push"
	"github.com/spf13/viper"
	"push_service/configs"
	"push_service/internal/models"
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
		unionIds []string
		content  string
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
				unionIds: []string{"on_8289e3180c41466f18c30f200fdc96e7"},
				content:  "case 1",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "case 2",
			ctx:  context.Background(),
			arg: arg{
				unionIds: []string{"on_8289e3180c41466f18c30f200fdc96e7"},
				content:  "case 2",
			},
			want: want{
				err: nil,
			},
		},
	}
	client := NewActionWithCtx(context.Background())
	client.SetMysql(mdb)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := &push.SendLarkTextByUnionIdsRequest{
				UnionIds: test.arg.unionIds,
				Content:  test.arg.content,
			}
			_, err := client.Deal(test.ctx, request)
			if err != test.want.err {
				t.Error(err)
				return
			}
			fmt.Println("发送飞书文本消息成功")
		})
	}
}
