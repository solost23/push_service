package base

import (
	"context"

	"github.com/gookit/slog"

	"gorm.io/gorm"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/solost23/protopb/gen/go/common"
	"push_service/configs"
)

type Action struct {
	ctx           context.Context
	mysqlConnect  *gorm.DB
	redisClient   *redis.Client
	kafkaProducer sarama.SyncProducer
	sl            *slog.SugaredLogger
	serverConfig  *configs.ServerConfig
	traceId       int64
	operator      int32
}

func (a *Action) SetContext(ctx context.Context) *Action {
	a.ctx = ctx
	return a
}

func (a *Action) SetHeader(header *common.RequestHeader) *Action {
	a.traceId = header.TraceId
	a.operator = header.OperatorId
	return a
}

func (a *Action) SetMysql(mysqlConn *gorm.DB) *Action {
	a.mysqlConnect = mysqlConn.WithContext(a.ctx)
	return a
}

func (a *Action) SetKafkaProducer(kafkaProducer sarama.SyncProducer) *Action {
	a.kafkaProducer = kafkaProducer
	return a
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

func (a *Action) SetSl(sl *slog.SugaredLogger) *Action {
	a.sl = sl
	return a
}

func (a *Action) GetSl() *slog.SugaredLogger {
	return a.sl
}

func (a *Action) SetServerConfig(serverConfig *configs.ServerConfig) *Action {
	a.serverConfig = serverConfig
	return a
}

func (a *Action) GetServerConfig() *configs.ServerConfig {
	return a.serverConfig
}

func (*Action) BuildError(code int32, msg string, header *common.RequestHeader) *common.ErrorInfo {
	if header == nil {
		header = new(common.RequestHeader)
	}
	return &common.ErrorInfo{
		Requester:  header.Requester,
		OperatorId: header.OperatorId,
		TraceId:    header.TraceId,
		Code:       code,
		Msg:        msg,
	}
}
