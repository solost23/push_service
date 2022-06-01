package service

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"my_grpc_frame/internal/service/create_hello_world"
	"my_grpc_frame/internal/service/delete_hello_world"
	"my_grpc_frame/internal/service/list_hello_world"
	"my_grpc_frame/internal/service/update_hello_world"
)

type HelloWorldService struct {
	mysqlConnect  *gorm.DB
	redisClient   *redis.Client
	kafkaProducer sarama.SyncProducer
	hello_world_service.UnimplementedMediaServiceServer
}

func NewHelloWorldService(mysqlConnect *gorm.DB, redisClient *redis.Client, kafkaProducer sarama.SyncProducer) *HelloWorldService {
	return &HelloWorldService{
		mysqlConnect:  mysqlConnect,
		redisClient:   redisClient,
		kafkaProducer: kafkaProducer,
	}
}

// 创建hello world
func (h *HelloWorldService) CreateHelloWorld(ctx context.Context, request *hello_world_service.CreateHelloWorldRequest) (reply *hello_world_service.CreateHelloWorldResponse, err error) {
	action := create_hello_world.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(h.mysqlConnect)
	action.SetkafkaProducer(h.kafkaProducer)
	return action.Deal(ctx, request)
}

// 删除hello world
func (h *HelloWorldService) DeleteHelloWorld(ctx context.Context, request *hello_world_service.DeleteHelloWorldRequest) (reply *hello_world_service.DeleteHelloWorldResponse, err error) {
	action := delete_hello_world.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(h.mysqlConnect)
	action.SetkafkaProducer(h.kafkaProducer)
	return action.Deal(ctx, request)
}

// 修改hello world
func (h *HelloWorldService) UpdateHelloWorld(ctx context.Context, request *Hello_world_service.UpdateHelloWorldRequest) (reply *hello_world_service.UpdateHelloWorldResponse, err error) {
	action := update_hello_world.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(h.mysqlConnect)
	action.SetkafkaProducer(h.kafkaProducer)
	return action.Deal(ctx, request)
}

func (h *HelloWorldService) ListHelloWorld(ctx context.Context, request *hello_world_service.ListHelloWorldRequest) (reply *hello_world_service.ListHelloWorldResponse, err error) {
	action := list_hello_world.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(h.mysqlConnect)
	action.SetkafkaProducer(h.kafkaProducer)
	return action.Deal(ctx, request)
}
