package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/mq"
	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.

	cartResult ,err := rpc.CartClient.GetCart(s.ctx,&cart.GetCartReq{UserId: req.UserId})
	if err!=nil{
		return nil,kerrors.NewGRPCBizStatusError(5005001,err.Error())
	}
	if cartResult ==nil || cartResult.Items == nil{
		return nil,kerrors.NewGRPCBizStatusError(5005001,"cart is Empty")
	}
	var total float32
	var oi  []*order.OrderItem

	for _ ,cartItem :=range cartResult.Items{
		productResp,resultErr :=rpc.ProductClient.GetProduct(s.ctx,&product.GetProductReq{
			Id:cartItem.ProductId,})

		if resultErr !=nil{
			return nil,resultErr
		}
		if productResp.Product ==nil{
			continue
		}
		p := productResp.Product.Price
		cost := p*float32(cartItem.Quantity)
		total+=cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	}
	var orderId string
	zipCodeInt, _ := strconv.Atoi(req.Address.ZipCode)
	orderResp,err := rpc.OrderClient.PlaceOrder(s.ctx,&order.PlaceOrderReq{
		UserId: req.UserId,
		Email: req.Email,
		UserCurrency: "USD",
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City: req.Address.City,
			ZipCode: int32(zipCodeInt),
			State: req.Address.State,
			Country: req.Address.Country,

		},
		OrderItems: oi,
	})
	if err!=nil{
		return nil,kerrors.NewGRPCBizStatusError(5005002,err.Error())
	}

	_,err =rpc.CartClient.EmptyCart(s.ctx,&cart.EmptyCartReq{UserId: req.UserId})
	if err!=nil{
		klog.Error(err.Error())
	}

	if orderResp !=nil || orderResp.Order!=nil{
		orderId = orderResp.GetOrder().OrderId
	}


	payReq := &payment.ChargeReq{
		UserId: req.UserId,
		OrderId:orderId,
		Amount:total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber: req.CreditCard.CreditCardNumber,
			CreditCardCvv: req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear: req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
	}
    
	paymentResult,err:=rpc.PaymentClient.Charge(s.ctx,payReq)
	if err != nil {
		err = fmt.Errorf("Charge.err:%v", err)
		return
	}
data, _ := proto.Marshal(&email.EmailReq{
    From:        "from@example.com",
    To:          req.Email,
    ContentType: "text/plain",
    Subject:     "You just created an order in CloudWeGo shop",
    Content:     "You just created an order in CloudWeGo shop",
})

msg := &nats.Msg{
    Subject: "email",
    Data:    data,
	Header: make(nats.Header),
}
otel.GetTextMapPropagator().Inject(s.ctx,propagation.HeaderCarrier(msg.Header))

err = mq.Nc.PublishMsg(msg)
if err != nil {
    fmt.Printf("Failed to publish message: %v", err)
} else {
    fmt.Println("Message published successfully")
}


	
	klog.Info(paymentResult)


	resp = &checkout.CheckoutResp{
		OrderId: orderId,
		TransactionId: paymentResult.TransactionId,
	}

	return resp,nil
}
