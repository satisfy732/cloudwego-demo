package notify

import (
	"fmt"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email"
	
)

type NoopEmail struct{}

func (e *NoopEmail) Send(req *email.EmailReq) error {
	fmt.Printf("Sending email: %v\n", req)
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}
