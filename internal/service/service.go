package service

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/solost23/protopb/gen/go/protos/push"
	"gorm.io/gorm"
	"push_service/internal/service/send_email"
	"push_service/internal/service/send_lark_post_by_union_ids"
	"push_service/internal/service/send_lark_text_by_union_ids"
)

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

func (p *PushService) SendEmail(ctx context.Context, request *push.SendEmailRequest) (reply *push.SendEmailResponse, err error) {
	action := send_email.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(p.mysqlConnect)
	action.SetkafkaProducer(p.kafkaProducer)
	return action.Deal(ctx, request)
}

func (p *PushService) SendLarkTextByUnionIds(ctx context.Context, request *push.SendLarkTextByUnionIdsRequest) (reply *push.SendLarkTextByUnionIdsResponse, err error) {
	action := send_lark_text_by_union_ids.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(p.mysqlConnect)
	action.SetkafkaProducer(p.kafkaProducer)
	return action.Deal(ctx, request)
}

func (p *PushService) SendLarkPostByUnionIds(ctx context.Context, request *push.SendLarkPostByUnionIdsRequest) (reply *push.SendLarkPostByUnionIdsResponse, err error) {
	action := send_lark_post_by_union_ids.NewActionWithCtx(ctx)
	action.SetHeader(request.Header)
	action.SetMysql(p.mysqlConnect)
	action.SetkafkaProducer(p.kafkaProducer)
	return action.Deal(ctx, request)
}
