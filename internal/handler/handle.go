package handler

import (
	"github.com/solost23/protopb/gen/go/protos/push"
	"push_service/internal/service"
)

func Init(config Config) (err error) {
	// 1.gRPC::push service
	push.RegisterPushServer(config.Server, service.NewPushService(config.MysqlConnect, config.RedisClient, config.KafkaProducer, config.Sl, config.ServerConfig))
	return
}
