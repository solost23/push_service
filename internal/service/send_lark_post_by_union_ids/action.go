package send_lark_post_by_union_ids

import (
	"context"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/solost23/protopb/gen/go/push"
	"push_service/internal/models"
	"push_service/internal/service/base"
	"push_service/pkg/utils"
)

type Action struct {
	base.Action
}

func NewActionWithCtx(ctx context.Context) *Action {
	a := &Action{}
	a.SetContext(ctx)
	return a
}

func (a *Action) Deal(_ context.Context, request *push.SendLarkPostByUnionIdsRequest) (reply *push.SendLarkPostByUnionIdsResponse, err error) {
	unionIds := request.GetUnionIds()
	var post map[string]interface{}
	err = jsoniter.UnmarshalFromString(string(request.GetContent().GetValue()), &post)
	if err != nil {
		return nil, err
	}
	err = utils.SendLarkMsg(map[string]interface{}{
		"msg_type":  "post",
		"union_ids": unionIds,
		"content": map[string]interface{}{
			"post": post,
		},
	}, a.GetServerConfig().LarkConfig[1].AppID, a.GetServerConfig().LarkConfig[1].AppSecret)
	if err != nil {
		return nil, err
	}

	go func() {
		// 记录日志
		if err = models.GInsert(a.GetMysqlConnect(), &models.LogLarkMsg{
			Type:     models.LarkMsgTypePost,
			UnionIds: strings.Join(unionIds, ","),
			Content:  string(request.GetContent().GetValue()),
		}); err != nil {
			a.GetSl().Error("存储失败: ", err.Error())
		}
	}()
	return reply, nil
}
