package handler

import (
	"github.com/solost23/my_interface/hello_world_service"

	"my_grpc_frame/internal/service"
)

func Init(config Config) (err error) {
	// 1.gRPC::hello world service
	hello_world_service.RegisterHelloWorldServiceServer(config.Server, service.NewHelloWorldService(config.MysqlConnect, config.RedisClient,
		config.KafkaProducer))
	return
}
