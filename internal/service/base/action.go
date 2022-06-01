package base

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
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
	return
}

func (a *Action) SetHeader(header *common.RequestHeader) {
	a.traceId = header.TraceId
	a.operator = header.OperatorUid
	return
}

func (a *Action) SetMysql(mysqlConn *gorm.DB) {
	a.mysqlConnect = mysqlconn.WithContext(a.ctx)
	return
}

func (a *Action) SetkafkaProducer(kafkaProducer sarama.SyncProducer) {
	a.kafkaProducer = kafkaProducer
	return
}

func (a *Action) GetTraceId() int64 {
	return a.traceId
}

func (a *Action) GetOperator() int32 {
	return a.operator
}

func (this *Action) GetMysqlConnect() *gorm.DB {
	return this.mysqlConnect
}

func (this *Action) GetRedisClient() *redis.Client {
	return this.redisClient
}

func (this *Action) GetKafkaProducer() sarama.SyncProducer {
	return this.kafkaProducer
}

func (*Action) BuildError(code error_code.ErrCode, msg string, header *common.RequestHeader) *common.ErrorInfo {
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

func (this *Action) ConvertProtoPage2DbPage(pager *common.RequestPager) (pageInfo models.PageInfo) {
	if pager == nil {
		pager = new(common.RequestPager)
	}
	var pageSize int32 = 10
	var pageIndex int32 = 0
	if pager.PageSize > 0 {
		pageSize = pager.PageSize
	}
	if pager.PageIndex > 0 {
		pageIndex = pager.PageIndex - 1
	}
	return models.PageInfo{
		NotPaging: false,
		Offset:    int(pageIndex * pageSize),
		Limit:     int(pageSize),
	}
}
