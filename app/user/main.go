// package main

// import (
// 	"log"
// 	"net"
// 	"time"

// 	"gomall_demo/app/user/biz/dal"
// 	"gomall_demo/app/user/conf"
// 	"gomall_demo/app/user/infra/rpc"
// 	"gomall_demo/common/mtl"
// 	"gomall_demo/common/serversuite"
// 	"gomall_demo/rpc_gen/kitex_gen/user/userservice"

// 	"github.com/cloudwego/kitex/pkg/klog"
// 	"github.com/cloudwego/kitex/pkg/rpcinfo"
// 	"github.com/cloudwego/kitex/pkg/transmeta"
// 	"github.com/cloudwego/kitex/server"
// 	"github.com/joho/godotenv"
// 	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
// 	"github.com/kitex-contrib/obs-opentelemetry/tracing"
// 	consul "github.com/kitex-contrib/registry-consul" // 复制过来的
// 	"go.uber.org/zap/zapcore"
// 	"gopkg.in/natefinch/lumberjack.v2"
// )

// var serviceName = conf.GetConf().Kitex.Service

// func main() {
// 	_ = godotenv.Load()
// 	mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])
// 	mtl.InitTracing(serviceName)
// 	mtl.InitLog(&lumberjack.Logger{
// 		Filename:   conf.GetConf().Kitex.LogFileName,
// 		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
// 		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
// 		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
// 	})
// 	dal.Init()
// 	opts := kitexInit()
// 	rpc.InitClient()

// 	svr := userservice.NewServer(new(UserServiceImpl), opts...)

// 	err := svr.Run()
// 	if err != nil {
// 		klog.Error(err.Error())
// 	}
// }

// func kitexInit() (opts []server.Option) {
// 	// address
// 	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
// 	log.Println("服务启动时监听的地址:", conf.GetConf().Kitex.Address)

// 	if err != nil {
// 		panic(err)
// 	}
// 	opts = append(opts, server.WithServiceAddr(addr))

// 	opts = append(opts,
// 		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
// 		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
// 		server.WithSuite(serversuite.CommonServerSuite{CurrentServiceName: serviceName}),
// 	)
// 	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
// 	if err != nil {
// 		klog.Fatal(err)
// 	}
// 	opts = append(opts, server.WithRegistry(r))

// 	opts = append(opts, server.WithSuite(tracing.NewServerSuite()))

// 	// klog
// 	logger := kitexlogrus.NewLogger()
// 	klog.SetLogger(logger)
// 	klog.SetLevel(conf.LogLevel())
// 	asyncWriter := &zapcore.BufferedWriteSyncer{
// 		WS: zapcore.AddSync(&lumberjack.Logger{
// 			Filename:   conf.GetConf().Kitex.LogFileName,
// 			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
// 			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
// 			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
// 		}),
// 		FlushInterval: time.Minute,
// 	}
// 	klog.SetOutput(asyncWriter)
// 	server.RegisterShutdownHook(func() {
// 		asyncWriter.Sync()
// 	})
// 	return
// }

package main

import (
	"net"
	"time"

	"gomall_demo/app/user/biz/dal"
	"gomall_demo/app/user/conf"
	"gomall_demo/app/user/infra/rpc"

	"gomall_demo/rpc_gen/kitex_gen/user/userservice"

	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv" // 复制过来的
	consul "github.com/kitex-contrib/registry-consul"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		klog.Error(err.Error())
	}

	dal.Init()
	opts := kitexInit()
	rpc.InitClient()

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
