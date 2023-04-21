package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/api/core/response"
	"github.com/larksuite/oapi-sdk-go/core"
	larkConfig "github.com/larksuite/oapi-sdk-go/core/config"
)

func SendLarkMsg(content map[string]interface{}, appID string, appSecret string) error {
	ret := make(map[string]interface{})

	req := request.NewRequestWithNative("/open-apis/message/v4/batch_send", "POST",
		request.AccessTokenTypeTenant, content, &ret)
	coreCtx := core.WrapContext(context.Background())
	// Send request

	cfg := getLarkConfig(appID, appSecret)
	err := api.Send(coreCtx, cfg, req)
	if err != nil {
		e := err.(*response.Error)
		return errors.New(fmt.Sprintf("%d:%s", e.Code, e.Msg))
	}

	return nil
}

func getLarkConfig(appID string, appSecret string) *larkConfig.Config {
	appSettings := core.NewInternalAppSettings(
		core.SetAppCredentials(appID, appSecret),
	)

	return core.NewConfig(core.DomainFeiShu,
		appSettings,
		core.SetLoggerLevel(core.LoggerLevelError))
}
