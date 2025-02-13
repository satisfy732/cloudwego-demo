// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	auth "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/auth"
	category "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/category"
	home "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/home"
	product "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	category.Register(r)

	product.Register(r)

	home.Register(r)

	auth.Register(r)
}
