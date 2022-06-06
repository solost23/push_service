package create_user

import (
	"context"

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

func (a *Action) Deal(_ context.Context, request *user_service.CreateUserRequest) (reply *user_service.CreateUserResponse, err error) {
	// 业务逻辑
	return reply, err
}
