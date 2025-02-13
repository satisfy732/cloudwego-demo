//主要业务逻辑

package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any,  error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	products,err :=rpc.ProductClient.ListProducts(h.Context,&product.ListProductsReq{})
	if err!=nil{
		return nil,err
	}
	return utils.H{
		"title":"Hot Sale",
		"items":products.Products,
	},nil
}
