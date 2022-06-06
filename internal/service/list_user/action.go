package list_user

import (
	"context"
	"github.com/solost23/go_interface/gen-go/common"
	"github.com/solost23/go_interface/gen-go/user_service"

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

func (a *Action) Deal(_ context.Context, request *user_service.ListUserRequest) (reply *user_service.ListUserResponse, err error) {
	// 业务逻辑
	reply.ErrorInfo = &common.ErrorInfo{}
	reply.User = &user_service.User{
		UserName: request.User.UserName,
		Password: request.User.Password,
	}
	return reply, err
}
