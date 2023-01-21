package base

import (
	"context"

	"gorm.io/gorm"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/solost23/protopb/gen/go/protos/common"
)

type Action struct {
	ctx           context.Context
	mysqlConnect  *gorm.DB
	redisClient   *redis.Client
	kafkaProducer sarama.SyncProducer
	traceId       int64
	operator      int32
}

func (a *Action) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func (a *Action) SetHeader(header *common.RequestHeader) {
	a.traceId = header.TraceId
	a.operator = header.OperatorUid
}

func (a *Action) SetMysql(mysqlConn *gorm.DB) {
	a.mysqlConnect = mysqlConn.WithContext(a.ctx)
}

func (a *Action) SetkafkaProducer(kafkaProducer sarama.SyncProducer) {
	a.kafkaProducer = kafkaProducer
}

func (a *Action) GetTraceId() int64 {
	return a.traceId
}

func (a *Action) GetOperator() int32 {
	return a.operator
}

func (a *Action) GetMysqlConnect() *gorm.DB {
	return a.mysqlConnect
}

func (a *Action) GetRedisClient() *redis.Client {
	return a.redisClient
}

func (a *Action) GetKafkaProducer() sarama.SyncProducer {
	return a.kafkaProducer
}

func (*Action) BuildError(code int32, msg string, header *common.RequestHeader) *common.ErrorInfo {
	if header == nil {
		header = new(common.RequestHeader)
	}
	return &common.ErrorInfo{
		Requester:   header.Requester,
		OperatorUid: header.OperatorUid,
		TraceId:     header.TraceId,
		Code:        code,
		Msg:         msg,
	}
}
