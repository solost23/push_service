package send_lark_post_by_union_ids

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/solost23/protopb/gen/go/push"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/anypb"
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
	dailyMsg := map[string]map[string]interface{}{
		"zh_cn": {
			"title": "工作日报",
			"content": [][]map[string]interface{}{
				{
					{
						"tag":  "text",
						"text": "忙碌了一天，开始整理今天的个人工作和收获吧!",
					},
				},
				{
					{
						"tag":  "a",
						"href": "https://www.baidu.com",
						"text": "点我填写今天的工作日报",
					},
				},
				{
					{
						"tag":  "text",
						"text": fmt.Sprintf("时间: %s", time.Now().Format("2006/01/02 15:04:05")),
					},
				},
			},
		},
	}
	weeklyMsg := map[string]map[string]interface{}{
		"zh_cn": {
			"title": "工作总结",
			"content": [][]map[string]interface{}{
				{
					{
						"tag":  "text",
						"text": "经过几天的工作，是时候对本周工作进行回顾和总结了!",
					},
				},
				{
					{
						"tag":  "a",
						"href": "https://www.baidu.com",
						"text": "点我填写工作总结",
					},
				},
				{
					{
						"tag":  "text",
						"text": fmt.Sprintf("时间: %s", time.Now().Format("2006/01/02 15:04:05")),
					},
				},
			},
		},
	}
	dailyStr, err := json.Marshal(dailyMsg)
	if err != nil {
		t.Error(err)
		return
	}
	weeklyStr, err := json.Marshal(weeklyMsg)
	if err != nil {
		t.Error(err)
		return
	}
	type arg struct {
		unionIds []string
		content  *anypb.Any
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
				content: &anypb.Any{
					Value: dailyStr,
				},
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
				content: &anypb.Any{
					Value: weeklyStr,
				},
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
			request := &push.SendLarkPostByUnionIdsRequest{
				UnionIds: test.arg.unionIds,
				Content:  test.arg.content,
			}
			_, err := client.Deal(test.ctx, request)
			if err != test.want.err {
				t.Error(err)
				return
			}
			fmt.Println("发送飞书富文本消息成功")
		})
	}
}
