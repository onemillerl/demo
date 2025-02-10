.PHONY: gen-user
	cwgo client --type RPC --service user --module gomall_demo/rpc_gen  -I ../idl  --idl ../idl/user.proto	
	cwgo server --type RPC --service user --module gomall_demo/app/user -pass "-use gomall_demo/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

.PHONY: gen-frontend
gen-frontend:
    cwgo server --type HTTP --idl ../../idl/frontend/hello.proto --service frontend --module github.com/cloudwego/biz-demo/gomall/app/frontend -I ../../idl/frontend/auth_page.proto
cwgo server -I ../../idl --type HTTP --service frontend --module  github.com/cloudwego/biz-demo/gomall/app/frontend --idl ../../idl/frontend/category_page.proto
cwgo server -I ../../idl --type HTTP --service frontend --module  gomall_demo/app/frontend --idl ../../idl/frontend/category_page.proto

gen-:
    cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend  --module github.com/cloudwego/biz-demo/gomall/app/frontend --I ../../idl/


.PHONY: gen-product
gen-product: 
	@cd rpc_gen && cwgo client --type RPC --service product --module gomall_demo/rpc_gen  -I ../idl  --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module gomall_demo/app/product --pass "-use gomall_demo/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/product.proto
cwgo server --type RPC --service product --module github.com/cloudwego/biz-demo/gomall/app/product --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto


.PHONY: gen-cart
gen-cart: 
	@cd rpc_gen && cwgo client --type RPC --service cart --module gomall_demo/rpc_gen  -I ../idl  --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module gomall_demo/app/cart --pass "-use gomall_demo/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/cart.proto


.PHONY: gen-payment
gen-payment: 
	@cd rpc_gen && cwgo client --type RPC --service payment --module gomall_demo/rpc_gen  -I ../idl  --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --service payment --module gomall_demo/app/payment --pass "-use gomall_demo/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/payment.proto


.PHONY: gen-checkout
gen-checkout: 
	@cd rpc_gen && cwgo client --type RPC --service checkout --module gomall_demo/rpc_gen  -I ../idl  --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --service checkout --module gomall_demo/app/checkout --pass "-use gomall_demo/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/checkout.proto

.PHONY: gen-order
gen-order: 
	@cd rpc_gen && cwgo client --type RPC --service order --module gomall_demo/rpc_gen  -I ../idl  --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module gomall_demo/app/order --pass "-use gomall_demo/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/order.proto

.PHONY: gen-email
gen-email: 
	@cd rpc_gen && cwgo client --type RPC --service email --module gomall_demo/rpc_gen  -I ../idl  --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC --service email --module gomall_demo/app/email --pass "-use gomall_demo/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/email.proto

.PHONY: gen-auth
gen-auth: 
	@cd rpc_gen && cwgo client --type RPC --service auth --module gomall_demo/rpc_gen  -I ../idl  --idl ../idl/auth.proto
	@cd app/auth && cwgo server --type RPC --service auth --module gomall_demo/app/auth --pass "-use gomall_demo/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/auth.proto
