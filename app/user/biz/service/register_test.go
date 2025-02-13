package service

import (
	"context"
	"testing"
	user "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal/mysql"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email :"demo@damin.com",
		Password :"122224fddfgg",
		ConfirmPassword:"122224fddfgg",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
