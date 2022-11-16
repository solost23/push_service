package send_email

import (
	"context"
	"errors"
	"github.com/go-gomail/gomail"
	"github.com/solost23/go_interface/gen_go/common"
	"github.com/solost23/go_interface/gen_go/push"
	"net/http"
	"strconv"

	"my_grpc_frame/internal/service/base"
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
		reply = &push.SendEmailResponse{
			ErrorInfo: &common.ErrorInfo{
				Code: http.StatusInternalServerError,
				Msg:  "发送邮件失败",
			},
		}
		return reply, err
	}
	reply = &push.SendEmailResponse{
		ErrorInfo: &common.ErrorInfo{
			Code: 0,
		},
	}
	return reply, err
}
