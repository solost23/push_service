package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	my_grpc_frame_log "github.com/solost23/tools/log"
	"github.com/spf13/viper"

	"my_grpc_frame/internal/server"
)

var (
	WebConfigPath = "configs/conf.yml"
	version       = "__BUILD_VERSION_"
	execDir       string
	st, v, V      bool
)

func main() {
	flag.StringVar(&execDir, "d", ".", "项目目录")
	flag.BoolVar(&v, "v", false, "查看版本号")
	flag.BoolVar(&V, "V", false, "查看版本号")
	flag.BoolVar(&st, "s", false, "项目状态")
	flag.Parse()
	if v || V {
		fmt.Println(version)
		return
	}
	// 运行
	InitConfig()
	InitLogger()
	server.Run()
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

func InitLogger() {
	// 默认已经正常load config到var cfg configs.Config
	// 使用自定义的log
	logger := my_grpc_frame_log.NewLogger(viper.GetString("log.runtime.path"))
	if logger == nil {
		fmt.Println("init logger failed")
		os.Exit(1)
	}
	ctxLogger := my_grpc_frame_log.NewLogger(viper.GetString("log.track.path"))
	if ctxLogger == nil {
		fmt.Println("init ctxLogger failed")
		os.Exit(1)
	}
}
