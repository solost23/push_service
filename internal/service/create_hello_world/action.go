package create_hello_world

import (
	"context"
	"my_grpc_frame/internal/service/base"

	"github.com/solost23/my_interface/hello_world_service"
)

type Action struct {
	base.Action
}

func NewActionWithCtx(ctx context.Context) *Action {
	a := &Action{}
	a.SetContext(ctx)
	return a
}

func (a *Action) Deal(_ context.Context, request *hello_world_service.CreateHelloWorldRequest) (reply *hello_world_service.CreateHelloWorldResponse, err error) {
	// 业务逻辑
	return reply, err
}
