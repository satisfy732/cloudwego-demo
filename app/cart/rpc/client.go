// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/cart/conf"
	cartutils "github.com/cloudwego/biz-demo/gomall/app/cart/utils"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
	err error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts :=[]client.Option{
       client.WithSuite(clientsuite.CommonClientSuite{
           CurrentServiceName: ServiceName,
		   RegistryAddr: RegistryAddr,
	   }),
	}
	ProductClient,err = productcatalogservice.NewClient("product",opts...)

	cartutils.MustHandleError(err)
}

