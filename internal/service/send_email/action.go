package send_email

import (
	"context"
	"net/http"

	"github.com/go-gomail/gomail"
	"github.com/solost23/protopb/gen/go/common"
	"github.com/solost23/protopb/gen/go/push"
	"push_service/internal/models"
	"push_service/internal/service/base"
)

type Action struct {
	base.Action
}

func NewActionWithCtx(ctx context.Context) *Action {
	a := &Action{}
	a.SetContext(ctx)
	return a
}

func (a *Action) Deal(_ context.Context, request *push.SendEmailRequest) (reply *push.SendEmailResponse, err error) {
	// 业务逻辑
	host := a.GetServerConfig().EmailConfig.Host
	port := a.GetServerConfig().EmailConfig.Port
	passwd := a.GetServerConfig().EmailConfig.Password
	name := a.GetServerConfig().EmailConfig.SendPersonName
	addr := a.GetServerConfig().EmailConfig.SendPersonAddr

	emailConf := request.GetEmail()
	m := gomail.NewMessage()
	m.SetAddressHeader("From", addr, name)
	m.SetHeader("To", m.FormatAddress(emailConf.GetAddr(), emailConf.GetName()))
	m.SetHeader("Subject", emailConf.GetTopic())
	m.SetBody(emailConf.GetContentType(), emailConf.GetContent())

	d := gomail.NewDialer(
		host,
		port,
		addr,
		passwd,
	)
	if err = d.DialAndSend(m); err != nil {
		_ = (&models.LogSendEmail{
			CreatorBase: models.CreatorBase{
				CreatorId: uint(request.GetHeader().GetOperatorId()),
			},
			Feature:       "邮件管理",
			OperationType: "发送邮件",
			Description:   request.GetEmail().GetContent(),
			Result:        false,
		}).Insert(a.GetMysqlConnect())
		reply = &push.SendEmailResponse{
			ErrorInfo: &common.ErrorInfo{
				Code: http.StatusInternalServerError,
				Msg:  "发送邮件失败",
			},
		}
		return reply, err
	}
	err = (&models.LogSendEmail{
		CreatorBase: models.CreatorBase{
			CreatorId: uint(request.GetHeader().GetOperatorId()),
		},
		Feature:       "邮件管理",
		OperationType: "发送邮件",
		Description:   request.GetEmail().GetContent(),
		Result:        true,
	}).Insert(a.GetMysqlConnect())
	if err != nil {
		return nil, err
	}
	reply = &push.SendEmailResponse{
		ErrorInfo: &common.ErrorInfo{
			Code: 0,
		},
	}
	return reply, err
}
