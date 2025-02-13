.PHONY:gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto &&cwgo server -I ../../idl --type RPC --module github.com/cloudwego/biz-demo/gomall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto
.PHONY:gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift &&cwgo server -I ../../idl --type RPC --module github.com/cloudwego/biz-demo/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift
.PHONY:gen-frontend
gen-frontend:
	@cd app/frontend &&cwgo server --type HTTP --idl ../../idl/frontend/product_page.proto --service frontend -module github.com/cloudwego/biz-demo/gomall/app/frontend -I ../../idl
.PHONY:gen-user
gen-user:
	@cd app/user && cwgo server --type RPC --idl ../../idl/user.proto --service user -module github.com/cloudwego/biz-demo/gomall/app/user -I ../../idl --pass "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen"
  
.PHONY:gen-product
gen-product:
	@cd app/product && cwgo server --type RPC --idl ../../idl/frontend/product.proto --service product --module github.com/cloudwego/biz-demo/gomall/app/product -I ../../idl --pass "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen"
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/frontend/product.proto --service product --module github.com/cloudwego/biz-demo/gomall/rpc_gen -I ../idl