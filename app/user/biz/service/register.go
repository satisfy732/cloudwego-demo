package service

import (
	"context"
	"errors"

	user "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"

	"golang.org/x/crypto/bcrypt"
	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/user/model"
	
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	if req.ConfirmPassword==""||req.Password== ""{
		return nil,errors.New("password is Empty")
	}
	if req.Email==""{
		return nil,errors.New("email is Empty")
	}
	

	if req.ConfirmPassword!=req.Password{
		return nil,errors.New("password not match")
	}
	// Finish your business logic.
	Passwordhashed,err:=bcrypt.GenerateFromPassword([]byte(req.Password),bcrypt.DefaultCost)
	if err!=nil{
		return nil,err 
	}
	newUser :=&model.User{
		Email:      req.Email,
		PasswordHashed: string(Passwordhashed),
	}
	err = model.Create(mysql.DB,newUser)
	if err!=nil{
		return nil,err 
	}
	return &user.RegisterResp{UserId:int32(newUser.ID)},nil
}
