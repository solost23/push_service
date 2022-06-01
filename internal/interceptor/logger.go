package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

func Logging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		return "myLogger", err
	}
}