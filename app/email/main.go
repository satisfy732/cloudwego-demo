package main

import (
	"net"
	"time"
	"context"
	

	"github.com/cloudwego/biz-demo/gomall/app/email/biz/consumer"
	"github.com/cloudwego/biz-demo/gomall/app/email/conf"
	"github.com/cloudwego/biz-demo/gomall/app/email/infra/mq"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email/emailservice"
	"github.com/cloudwego/kitex/pkg/klog"
	
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"github.com/cloudwego/biz-demo/gomall/common/mtl"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/cloudwego/biz-demo/gomall/common/serversuite"
)

var(
	ServiceName = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)
func main() {
	mq.Init()
	p:=mtl.InitTracing(ServiceName)
	defer p.Shutdown(context.Background())
	mtl.InitMetric(ServiceName,conf.GetConf().Kitex.MetricsPort,RegistryAddr)
	consumer.Init()
	opts := kitexInit()

	svr := emailservice.NewServer(new(EmailServiceImpl), opts...)

	err := svr.Run()
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
