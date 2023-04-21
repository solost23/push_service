package main

import (
	"flag"
	"fmt"
	_ "github.com/spf13/viper/remote"
	"os"
	"path"

	push_service_log "github.com/solost23/tools/log"
	"github.com/spf13/viper"

	"push_service/configs"
	"push_service/internal/server"
	"push_service/pkg/helper"
)

var (
	WebConfigPath = "configs/conf.yml"
	WebLogPath    = "logs"
	version       = "__BUILD_VERSION_"
	execDir       string
	provider      string
	st, v, V      bool
)

func main() {
	flag.StringVar(&execDir, "d", ".", "项目目录")
	flag.StringVar(&provider, "p", "consul", "项目配置提供者")
	flag.BoolVar(&v, "v", false, "查看版本号")
	flag.BoolVar(&V, "V", false, "查看版本号")
	flag.BoolVar(&st, "s", false, "项目状态")
	flag.Parse()
	if v || V {
		fmt.Println(version)
		return
	}
	// 运行
	//InitConfig()
	serverConfig, err := InitConfigFromConsul()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	sl, err := helper.InitLogger(execDir, WebLogPath, serverConfig.Mode)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	server.NewServer(serverConfig, sl).Run()
}

func InitConfig() {
	// viper.SetConfigFile(WebConfigPath)
	// viper.AddConfigPath(".")
	configPath := path.Join(execDir, WebConfigPath)
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("未找到配置文件，当前path:", configPath)
		panic(err)
	}
}

func InitConfigFromConsul() (serverConfig *configs.ServerConfig, err error) {
	serverConfig = new(configs.ServerConfig)
	v := viper.New()
	// 通过项目配置读取基本信息
	v.SetConfigFile(path.Join(execDir, WebConfigPath))
	if err = v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err = v.Unmarshal(serverConfig); err != nil {
		return nil, err
	}

	// 从配置中心读取配置
	err = v.AddRemoteProvider(provider,
		fmt.Sprintf("%s:%d", serverConfig.ConsulConfig.Host, serverConfig.ConsulConfig.Port),
		serverConfig.ConfigPath)
	if err != nil {
		return nil, err
	}

	v.SetConfigType("YAML")

	if err = v.ReadRemoteConfig(); err != nil {
		return nil, err
	}

	if err = v.Unmarshal(serverConfig); err != nil {
		return nil, err
	}
	return serverConfig, nil

}

func InitLogger() {
	// 默认已经正常load config到var cfg configs.Config
	// 使用自定义的log
	logger := push_service_log.NewLogger(viper.GetString("log.runtime.path"))
	if logger == nil {
		fmt.Println("init logger failed")
		os.Exit(1)
	}
	ctxLogger := push_service_log.NewLogger(viper.GetString("log.track.path"))
	if ctxLogger == nil {
		fmt.Println("init ctxLogger failed")
		os.Exit(1)
	}
}
