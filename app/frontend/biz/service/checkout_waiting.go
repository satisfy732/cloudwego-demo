package service

import (
	"context"

	checkout "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	rpcpayment "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment"
	rpccheckout "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string] any, err error) {
     userId := frontendutils.GetUserIdFromCtx(h.Context)
	 _,err = rpc.CheckoutClient.Checkout(h.Context,&rpccheckout.CheckoutReq{
		UserId: uint32(userId),
		Email: req.Email,
		Firstname: req.Firstname,
		Lastname: req.Lastname,
		Address: &rpccheckout.Address{
			Country: req.Country,
			City: req.City,
			ZipCode: req.Zipcode,
			State: req.Province,
			StreetAddress: req.Street,
		},
		CreditCard:&rpcpayment.CreditCardInfo{
			CreditCardNumber: req.CardNum,
			CreditCardCvv: req.Cvv,
			CreditCardExpirationYear: req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
		} ,
	 })
	 if err!=nil{
		return nil,err
	 }


	return utils.H{
		"title":"Waiting",
		"redirect":"/checkout/result",
	},nil
}
