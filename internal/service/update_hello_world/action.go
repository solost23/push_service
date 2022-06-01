package update_hello_world

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

func (a *Action) Deal(_ context.Context, request *hello_world_service.UpdateHelloWorldRequest) (reply *hello_world_service.UpdateHelloWorldResponse, err error) {
	// 业务逻辑
	return reply, err
}
