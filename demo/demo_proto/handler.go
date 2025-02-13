package main

import (
	"context"
	pbapi "github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/service"
)

// EchoserviceImpl implements the last service interface defined in the IDL.
type EchoserviceImpl struct{}

// Echo implements the EchoserviceImpl interface.
func (s *EchoserviceImpl) Echo(ctx context.Context, req *pbapi.Request) (resp *pbapi.Reponse, err error) {
	resp, err = service.NewEchoService(ctx).Run(req)

	return resp, err
}
