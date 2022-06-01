package delete_hello_world

import (
	"context"
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

func (a *Action) Deal(_ context.Context, request *hello_world_service.DeleteHelloWorldRequest) (reply *hello_world_service.DeleteHelloWorldResponse, err error) {
	// 业务逻辑
	return reply, err
}
