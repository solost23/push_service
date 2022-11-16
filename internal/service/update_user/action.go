package update_user

import (
	"context"

	"github.com/solost23/go_interface/gen_go/common"
	"github.com/solost23/go_interface/gen_go/user_service"

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

func (a *Action) Deal(_ context.Context, request *user_service.UpdateUserRequest) (reply *user_service.UpdateUserResponse, err error) {
	// 业务逻辑
	reply = &user_service.UpdateUserResponse{
		ErrorInfo: &common.ErrorInfo{
			Code: 0,
		},
		User: &user_service.User{
			UserName: request.User.UserName,
			Password: request.User.Password,
		},
	}
	return reply, err
}
