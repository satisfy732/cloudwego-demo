package rpc

import (
	"fmt"
	"sync"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/kitex-contrib/config-consul/consul"
	consulclient "github.com/kitex-contrib/config-consul/client"

	"context"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
)
var (
	UserClient userservice.Client
	ProductClient productcatalogservice.Client
	CartClient cartservice.Client
	OrderClient orderservice.Client
	CheckoutClient checkoutservice.Client
	once sync.Once
	ServiceName =frontendUtils.ServiceName
	RegistryAddr = conf.GetConf().Hertz.RegistryAddr
	err error
)

func Init(){
   once.Do(func(){
	iniUserClient()
	initProductClient()
	initCartClient()
	initCheckoutClient()
	iniOrderClient()
   })
}
func iniUserClient(){
	
	UserClient,err = userservice.NewClient("user",client.WithSuite(clientsuite.CommonClientSuite{
       CurrentServiceName: ServiceName,
	   RegistryAddr: RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		var res string = circuitbreak.RPCInfo2Key(ri)
		fmt.Printf("%s\n",res)
		return res;
	})
	//res和UpdateServiceCBConfig的第一个参数对应
	
	cbs.UpdateServiceCBConfig("frontend/product/ListProducts",
	    circuitbreak.CBConfig{Enable: true,ErrRate:0.5,MinSample:2 },
	)
	consulClient,err:=consul.NewClient(consul.Options{})
	if err != nil {
		panic(err)
	}
	ProductClient,err = productcatalogservice.NewClient("product",client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr: RegistryAddr,
	 }),client.WithCircuitBreaker(cbs),client.WithFallback(
		fallback.NewFallbackPolicy(
			fallback.UnwrapHelper(
				func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
					if err == nil {
						return resp, nil
					}
					methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
					fmt.Printf("fallback method name:%s\n",methodName)
					// ListProducts接口不做降级处理
					if methodName == "ListProducts" {
						return resp, err
					}
					return &product.ListProductsResp{
						Products: []*product.Product{
							{
								Price:       6.6,
								Id:          3,
								Picture:     "/static/image/t-shirt.jpeg",
								Name:        "T-Shirt",
								Description: "CloudWeGo T-Shirt",
							},
						},
	               },nil
				}),
		),
	 ),
    client.WithSuite(consulclient.NewSuite("product",ServiceName,consulClient)),
	)
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	CartClient,err = cartservice.NewClient("cart",client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr: RegistryAddr,
	 }))
	 frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient,err = checkoutservice.NewClient("checkout",client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr: RegistryAddr,
	 }))
	 frontendUtils.MustHandleError(err)
}
func iniOrderClient(){
	OrderClient,err = orderservice.NewClient("order",client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr: RegistryAddr,
	 }))
	 frontendUtils.MustHandleError(err)
}