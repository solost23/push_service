package handler

import (
	"github.com/solost23/go_interface/gen-go/user_service"

	"my_grpc_frame/internal/service"
)

func Init(config Config) (err error) {
	// 1.gRPC::user service
	user_service.RegisterUserServiceServer(config.Server, service.NewUserService(config.MysqlConnect, config.RedisClient,
		config.KafkaProducer))
	return
}
