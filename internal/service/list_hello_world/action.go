package list_hello_world

import (
	"context"

	"github.com/solost23/my_interface/hello_world_service"

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

func (a *Action) Deal(_ context.Context, request *hello_world_service.ListHelloWorldRequest) (reply *hello_world_service.ListHelloWorldResponse, err error) {
	// 业务逻辑
	return reply, err
}
