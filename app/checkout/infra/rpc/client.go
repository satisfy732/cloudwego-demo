package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/cart/conf"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"

)

var (
	CartClient cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
    OrderClient orderservice.Client
	ServiceName = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
	err error
	once   sync.Once

)

func InitClient(){
	once.Do(func ()  {
		InitCartClient()
		InitPaymentClient()
		InitProductClient()
		InitOrderClient()
	})
}

func InitCartClient(){
	opts :=[]client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr: RegistryAddr,
		}),
	 }
	 CartClient,err =cartservice.NewClient("cart",opts...)
 
	 if err!=nil{
		 panic(err)
	 }

}
func InitProductClient() {
	opts :=[]client.Option{
       client.WithSuite(clientsuite.CommonClientSuite{
           CurrentServiceName: ServiceName,
		   RegistryAddr: RegistryAddr,
	   }),
	}
	ProductClient,err = productcatalogservice.NewClient("product",opts...)

	if err!=nil{
		panic(err)
	}
	
}
func InitPaymentClient(){
	opts :=[]client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr: RegistryAddr,
		}),
	 }
	PaymentClient,err = paymentservice.NewClient("payment",opts ...)
	if err!=nil{
		panic(err)
	}
	
}
func InitOrderClient(){
	opts :=[]client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr: RegistryAddr,
		}),
	 }
	OrderClient,err = orderservice.NewClient("order",opts ...)
	if err!=nil{
		panic(err)
	 }
	
}