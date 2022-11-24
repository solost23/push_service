package handler

import (
	"github.com/solost23/go_interface/gen_go/push"
	"push_service/internal/service"
)

func Init(config Config) (err error) {
	// 1.gRPC::push service
	push.RegisterPushServer(config.Server, service.NewPushService(config.MysqlConnect, config.RedisClient, config.KafkaProducer))
	return
}
