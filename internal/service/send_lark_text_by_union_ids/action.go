package send_lark_text_by_union_ids

import (
	"context"
	"github.com/solost23/protopb/gen/go/protos/push"
	"push_service/internal/models"
	"push_service/internal/service/base"
	"push_service/pkg/utils"
	"strings"
)

type Action struct {
	base.Action
}

func NewActionWithCtx(ctx context.Context) *Action {
	a := &Action{}
	a.SetContext(ctx)
	return a
}

func (a *Action) Deal(_ context.Context, request *push.SendLarkTextByUnionIdsRequest) (reply *push.SendLarkTextByUnionIdsResponse, err error) {
	unionIds := request.GetUnionIds()
	content := request.GetContent()
	err = utils.SendLarkMsg(map[string]interface{}{
		"msg_type":  "text",
		"union_ids": unionIds,
		"content": map[string]string{
			"text": content,
		},
	}, a.GetServerConfig().LarkConfig[1].AppID, a.GetServerConfig().LarkConfig[1].AppSecret)
	if err != nil {
		return nil, err
	}

	go func() {
		// 记录日志
		if err = models.GInsert(a.GetMysqlConnect(), &models.LogLarkMsg{
			Type:     models.LarkMsgTypeText,
			UnionIds: strings.Join(unionIds, ","),
			Content:  content,
		}); err != nil {
			a.GetSl().Error("存储失败: ", err.Error())
		}
	}()
	return reply, nil
}
