package send_email

import (
	"context"
	"errors"
	"github.com/go-gomail/gomail"
	"github.com/solost23/protopb/gen/go/protos/common"
	"github.com/solost23/protopb/gen/go/protos/push"
	"net/http"
	"push_service/internal/models"
	"strconv"

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
	emailConf := request.GetEmail()
	m := gomail.NewMessage()
	m.SetAddressHeader("From", emailConf.GetSendPersonAddr(), emailConf.GetSendPersonName())
	m.SetHeader("To", m.FormatAddress(emailConf.GetAddr(), emailConf.GetName()))
	m.SetHeader("Subject", emailConf.GetTopic())
	m.SetBody(emailConf.GetContentType(), emailConf.GetContent())
	port, err := strconv.Atoi(emailConf.GetPort())
	if err != nil {
		_ = (&models.LogSendEmail{
			CreatorBase: models.CreatorBase{
				CreatorId: uint(request.GetHeader().GetOperatorUid()),
			},
			Feature:       "邮件管理",
			OperationType: "发送邮件",
			Description:   request.GetEmail().GetContent(),
			Result:        false,
		}).Insert(a.GetMysqlConnect())
		reply = &push.SendEmailResponse{
			ErrorInfo: &common.ErrorInfo{
				Code: http.StatusBadRequest,
				Msg:  "端口参数错误",
			},
		}
		return reply, errors.New("端口参数错误")
	}
	d := gomail.NewDialer(
		emailConf.GetHost(),
		port,
		emailConf.GetSendPersonAddr(),
		emailConf.GetPassword(),
	)
	if err = d.DialAndSend(m); err != nil {
		_ = (&models.LogSendEmail{
			CreatorBase: models.CreatorBase{
				CreatorId: uint(request.GetHeader().GetOperatorUid()),
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
			CreatorId: uint(request.GetHeader().GetOperatorUid()),
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
