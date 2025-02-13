package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/middleware"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	consul "github.com/kitex-contrib/registry-consul"

	"github.com/cloudwego/kitex/transport"
)

func main() {

    r, err := consul.NewConsulResolver("127.0.0.1:8500")
    if err != nil {
        log.Fatal(err)
    }
	client,err :=echoservice.NewClient("demo_proto",client.WithResolver(r),
                client.WithTransportProtocol(transport.GRPC),
                client.WithMetaHandler(transmeta.ClientHTTP2Handler),
                client.WithMiddleware(middleware.Middleware),
    )
   
    if err != nil {
        log.Fatal(err)
    }
	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")
    res,err := client.Echo(ctx,&pbapi.Request{Message:"hello"})
	var bizErr *kerrors.GRPCBizStatusError
	if(err != nil){
		ok := errors.As(err,&bizErr)
		if ok{
			fmt.Printf("%#v",bizErr)
		}
		log.Fatal(err)
	}
	fmt.Printf("response:%v\n",res)
}
