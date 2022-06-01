package handler

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/go-redis/redis"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Config struct {
	Server        *grpc.Server
	MysqlConnect  *gorm.DB
	RedisClient   *redis.Client
	KafkaConsumer *cluster.Consumer
	KafkaProducer sarama.SyncProducer
}
