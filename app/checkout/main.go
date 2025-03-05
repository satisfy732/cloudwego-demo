package main

import (
	
	"net"
	"time"
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/biz/dal"
	"github.com/cloudwego/biz-demo/gomall/app/checkout/conf"
	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/mq"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/cloudwego/biz-demo/gomall/common/mtl"
	"github.com/cloudwego/biz-demo/gomall/common/serversuite"
	
)
var(
	ServiceName = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)
func main() {
	_ = godotenv.Load()
	mtl.InitMetric(ServiceName,conf.GetConf().Kitex.MetricsPort,RegistryAddr)
	p:=mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background())
	rpc.InitClient()
	
	mq.Init()
	opts := kitexInit()


	svr := checkoutservice.NewServer(new(CheckoutServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	dal.Init()
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	
	opts = append(opts, server.WithServiceAddr(addr),server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr: RegistryAddr,
	}))
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
