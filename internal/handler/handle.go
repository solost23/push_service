package handler

import (
	"my_grpc_frame/internal/service"

	"github.com/solost23/my_interface/hello_world_service"
)

func Init(config Config) (err error) {
	// 1.gRPC::hello world service
	hello_world_service.RegisterHelloWorldServiceServer(config.Server, service.NewHelloWorldService(config.MysqlConnect, config.RedisClient,
		config.KafkaProducer))
	return
}
