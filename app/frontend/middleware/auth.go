package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	
	"github.com/hertz-contrib/sessions"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
)


func GlobalAuth() app.HandlerFunc{
   return func (ctx context.Context,c *app.RequestContext){
	s:=sessions.Default(c)
	ctx = context.WithValue(ctx,frontendUtils.SessionsUserId,s.Get("user_id"))
	c.Next(ctx)
   }
}

func Auth() app.HandlerFunc {
	return func (ctx context.Context,c *app.RequestContext){
		s:=sessions.Default(c)
		
		userId := s.Get("user_id")
		if userId ==nil{
			c.Redirect(302,[]byte ("/sign-in?next="+c.FullPath()))
			c.Abort()
			return 
		}
		c.Next(ctx)
	
	   }
}