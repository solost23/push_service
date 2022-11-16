package handler

import (
	"github.com/solost23/go_interface/gen_go/push"
	"github.com/solost23/go_interface/gen_go/user_service"

	"my_grpc_frame/internal/service"
)

func Init(config Config) (err error) {
	// 1.gRPC::user service
	user_service.RegisterUserServiceServer(config.Server, service.NewUserService(config.MysqlConnect, config.RedisClient,
		config.KafkaProducer))

	// 2.gRPC::push service
	push.RegisterPushServer(config.Server, service.NewPushService(config.MysqlConnect, config.RedisClient, config.KafkaProducer))
	return
}
