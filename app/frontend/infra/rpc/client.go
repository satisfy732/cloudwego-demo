package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	consul "github.com/kitex-contrib/registry-consul"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	
)
var (
	UserClient userservice.Client
	ProductClient productcatalogservice.Client
	once sync.Once
)

func Init(){
   once.Do(func(){
	iniUserClient()
	initProductClient()
   })
}
func iniUserClient(){
	var opts []client.Option
	r,err:=consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts,client.WithResolver(r))
	
	UserClient,err = userservice.NewClient("user",opts...)
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r,err:=consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts,client.WithResolver(r))
	ProductClient,err = productcatalogservice.NewClient("product",opts...)

	frontendUtils.MustHandleError(err)
}