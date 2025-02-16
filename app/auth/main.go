// package main

// import (
// 	"log"
// 	"net"
// 	"time"

// 	"gomall_demo/app/auth/biz/dal"
// 	"gomall_demo/app/auth/conf"
// 	"gomall_demo/common/mtl"
// 	"gomall_demo/common/serversuite"
// 	"gomall_demo/rpc_gen/kitex_gen/auth/authservice"

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

// 	svr := authservice.NewServer(new(AuthServiceImpl), opts...)

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
	"strings"

	"gomall_demo/app/auth/biz/dal"
	"gomall_demo/app/auth/conf"
	"gomall_demo/app/auth/infra/rpc"
	"gomall_demo/common/mtl"
	"gomall_demo/common/serversuite"
	"gomall_demo/common/utils"
	"gomall_demo/rpc_gen/kitex_gen/auth/authservice"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv" // 复制过来的
	"gopkg.in/natefinch/lumberjack.v2"
)

var serviceName = conf.GetConf().Kitex.Service

func main() {
	err := godotenv.Load()
	mtl.InitMetric(serviceName, conf.GetConf().Kitex.MetricsPort, conf.GetConf().Registry.RegistryAddress[0])
	mtl.InitTracing(serviceName)
	mtl.InitLog(&lumberjack.Logger{
		Filename:   conf.GetConf().Kitex.LogFileName,
		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
	})
	if err != nil {
		klog.Error(err.Error())
	}

	dal.Init()
	opts := kitexInit()
	rpc.InitClient()

	svr := authservice.NewServer(new(AuthServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	address := conf.GetConf().Kitex.Address
	if strings.HasPrefix(address, ":") {
		localIp := utils.MustGetLocalIPv4()
		address = localIp + address
	}
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{CurrentServiceName: serviceName, RegistryAddr: conf.GetConf().Registry.RegistryAddress[0]}))
	return
}
