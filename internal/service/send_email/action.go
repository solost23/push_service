package send_email

import (
	"context"
	"github.com/solost23/go_interface/gen_go/common"
	"github.com/solost23/go_interface/gen_go/push"

	"my_grpc_frame/internal/service/base"
)

type Action struct {
	base.Action
}

func NewActionWithCtx(ctx context.Context) *Action {
	a := &Action{}
	a.SetContext(ctx)
	return a
}

func (a *Action) Deal(ctx context.Context, request *push.SendEmailRequest) (reply *push.SendEmailResponse, err error) {
	// 业务逻辑

	reply = &push.SendEmailResponse{
		ErrorInfo: &common.ErrorInfo{
			Code: 0,
		},
	}
	return reply, err
}
