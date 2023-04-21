package handler

import (
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/gookit/slog"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"push_service/configs"
)

type Config struct {
	Server        *grpc.Server
	Sl            *slog.SugaredLogger
	MysqlConnect  *gorm.DB
	RedisClient   *redis.Client
	KafkaProducer sarama.SyncProducer
	ServerConfig  *configs.ServerConfig
}
