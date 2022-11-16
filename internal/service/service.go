package service

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/solost23/go_interface/gen_go/push"
	"github.com/solost23/go_interface/gen_go/user_service"
	"gorm.io/gorm"
	"my_grpc_frame/internal/service/send_email"

	"my_grpc_frame/internal/service/create_user"
	"my_grpc_frame/internal/service/delete_user"
	"my_grpc_frame/internal/service/list_user"
	"my_grpc_frame/internal/service/update_user"
)

type UserService struct {
	mysqlConnect  *gorm.DB
	redisClient   *redis.Client
	kafkaProducer sarama.SyncProducer
	user_service.UnimplementedUserServiceServer
}

func NewUserService(mysqlConnect *gorm.DB, redisClient *redis.Client, kafkaProducer sarama.SyncProducer) *UserService {
	return &UserService{
		mysqlConnect:  mysqlConnect,
		redisClient:   redisClient,
		kafkaProducer: kafkaProducer,
	}
}

// 创建User
func (h *UserService) CreateUser(ctx context.Context, request *user_service.CreateUserRequest) (reply *user_service.CreateUserResponse, err error) {
	action := create_user.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(h.mysqlConnect)
	action.SetkafkaProducer(h.kafkaProducer)
	return action.Deal(ctx, request)
}

// 删除User
func (h *UserService) DeleteUser(ctx context.Context, request *user_service.DeleteUserRequest) (reply *user_service.DeleteUserResponse, err error) {
	action := delete_user.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(h.mysqlConnect)
	action.SetkafkaProducer(h.kafkaProducer)
	return action.Deal(ctx, request)
}

// 修改User
func (h *UserService) UpdateUser(ctx context.Context, request *user_service.UpdateUserRequest) (reply *user_service.UpdateUserResponse, err error) {
	action := update_user.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(h.mysqlConnect)
	action.SetkafkaProducer(h.kafkaProducer)
	return action.Deal(ctx, request)
}

// 展示User
func (h *UserService) ListUser(ctx context.Context, request *user_service.ListUserRequest) (reply *user_service.ListUserResponse, err error) {
	action := list_user.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(h.mysqlConnect)
	action.SetkafkaProducer(h.kafkaProducer)
	return action.Deal(ctx, request)
}

// 推送服务
type PushService struct {
	mysqlConnect  *gorm.DB
	redisClient   *redis.Client
	kafkaProducer sarama.SyncProducer
	push.UnimplementedPushServer
}

func NewPushService(mysqlConnect *gorm.DB, redisClient *redis.Client, kafkaProducer sarama.SyncProducer) *PushService {
	return &PushService{
		mysqlConnect:  mysqlConnect,
		redisClient:   redisClient,
		kafkaProducer: kafkaProducer,
	}
}

func (p *PushService) SendMail(ctx context.Context, request *push.SendEmailRequest) (reply *push.SendEmailResponse, err error) {
	action := send_email.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(p.mysqlConnect)
	action.SetkafkaProducer(p.kafkaProducer)
	return action.Deal(ctx, request)
}
