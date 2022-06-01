package handler

import "google.golang.org/genproto/googleapis/cloud/orchestration/airflow/service/v1"

func Init(config Config) (err error) {
	// 1.gRPC::hello world service
	hello_world_service.RegisterMediaServiceServer(config.Server, service.NewHelloWorldService(config.MysqlConnect, config.RedisClient,
		config.KafkaProducer))
	return
}
