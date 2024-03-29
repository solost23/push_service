package server

import (
	"fmt"
	"github.com/gookit/slog"
	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	"push_service/configs"
	"push_service/internal/handler"
	"push_service/internal/interceptor"
	"push_service/internal/models"
	"push_service/pkg/helper"
)

type Server struct {
	serverConfig *configs.ServerConfig
	// 日志句柄
	sl *slog.SugaredLogger
}

func NewServer(serverConfig *configs.ServerConfig, sl *slog.SugaredLogger) *Server {
	return &Server{
		serverConfig: serverConfig,
		sl:           sl,
	}
}

func (s *Server) Run() {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptor.Context(),
			interceptor.Recovery(),
			interceptor.Logging(),
		)))
	mdb, err := models.InitMysql(s.serverConfig.MySQLConfig)
	must(err)
	// 初始化 handler
	err = handler.Init(handler.Config{
		Server:       server,
		MysqlConnect: mdb,
		RedisClient:  nil,
		Sl:           s.sl,
		ServerConfig: s.serverConfig,
	})
	must(err)
	// 随机获取ip 地址和 未占用端口
	ip, err := helper.GetInternalIP()
	must(err)
	port, err := helper.GetFreePort()
	must(err)

	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	must(err)

	// 注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 服务注册
	cfg := api.DefaultConfig()
	cfg.Address = s.serverConfig.ConsulConfig.Host
	cfg.Address = fmt.Sprintf("%s:%d", s.serverConfig.ConsulConfig.Host, s.serverConfig.ConsulConfig.Port)

	client, err := api.NewClient(cfg)
	must(err)

	// 生成检查对象
	check := &api.AgentServiceCheck{
		Interval:                       "5s",
		Timeout:                        "5s",
		GRPC:                           fmt.Sprintf("%s:%d", ip, port),
		DeregisterCriticalServiceAfter: "10s",
	}

	// 生成注册对象
	serviceId := uuid.NewV4().String()
	registration := new(api.AgentServiceRegistration)
	registration.Name = s.serverConfig.Name
	registration.ID = serviceId
	registration.Port = port
	registration.Tags = strings.Split(s.serverConfig.Name, "_")
	registration.Address = ip
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	must(err)

	// 启动gRPC server
	go func() {
		defer RecoverGRoutine("GRpc")
		err = server.Serve(listen)
		if err != nil {
			panic(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			// 服务停止
			server.Stop()
			// 服务注销
			if err = client.Agent().ServiceDeregister(serviceId); err != nil {
				fmt.Println("服务注销失败")
			}
			fmt.Println("服务注销成功")
			// kafka.ConsumerClient.Stop()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func RecoverGRoutine(qid interface{}) {
	if err := recover(); err != nil {
		buf := make([]byte, 1<<16)
		runtime.Stack(buf, true)
		fmt.Println(fmt.Sprintf("[%v] qid:%v, panic:%v\n, runtime:%v", time.Now().Format("2006-01-02 15:04:05"), qid, err, string(buf)))
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
