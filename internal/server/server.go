package server

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"my_grpc_frame/internal/handler"
	"my_grpc_frame/internal/interceptor"
	"my_grpc_frame/internal/models"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func Run() {
	port := viper.GetInt("params.service_port")
	server, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
		return
	}
	gRPCServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		interceptor.Context(),
		interceptor.Recovery(),
		interceptor.Logging())))
	// 初始化 handler
	err = handler.Init(handler.Config{
		Server:       gRPCServer,
		MysqlConnect: models.NewMysqlConnect(),
		RedisClient:  models.NewRedisConnect("hello world"),
	})
	if err != nil {
		panic(err)
		return
	}
	// 启动gRPC server
	go func() {
		defer RecoverGRoutine("GRpc")
		err = gRPCServer.Serve(server)
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
			gRPCServer.Stop()
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
