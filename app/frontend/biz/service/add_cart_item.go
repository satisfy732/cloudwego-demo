package service

import (
	"context"

	cart "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/cart"
	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	rpccart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	
	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()d
	// todo edit your code
	_,err = rpc.CartClient.AddItem(h.Context,&rpccart.AddItemReq{
		UserId:uint32(frontendutils.GetUserIdFromCtx(h.Context)),
		Item:&rpccart.CartItem{
			ProductId:req.ProductId,
			Quantity:req.ProductNum,
		},

	})
	if err!=nil{
		return nil,err
	}

	return 
}
