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
	"github.com/spf13/viper"
)

func SendLarkMsg(content map[string]interface{}, larkAppId uint) error {
	ret := make(map[string]interface{})

	req := request.NewRequestWithNative("/open-apis/message/v4/batch_send", "POST",
		request.AccessTokenTypeTenant, content, &ret)
	coreCtx := core.WrapContext(context.Background())
	// Send request

	cfg := getLarkConfig(larkAppId)
	err := api.Send(coreCtx, cfg, req)
	if err != nil {
		e := err.(*response.Error)
		return errors.New(fmt.Sprintf("%d:%s", e.Code, e.Msg))
	}

	return nil
}

func getLarkConfig(id uint) *larkConfig.Config {
	appId := viper.GetString(fmt.Sprintf("lark.%d.app_id", id))
	appSecret := viper.GetString(fmt.Sprintf("lark.%d.app_secret", id))
	appSettings := core.NewInternalAppSettings(
		core.SetAppCredentials(appId, appSecret),
	)

	return core.NewConfig(core.DomainFeiShu,
		appSettings,
		core.SetLoggerLevel(core.LoggerLevelError))
}
