package service

import (
	"gomall_demo/rpc_gen/kitex_gen/cart/cartservice"

	cartutils "gomall_demo/app/cart/utils"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var cartClient cartservice.Client

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	cartutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	cartClient, err = cartservice.NewClient("cart", opts...)
	cartutils.MustHandleError(err)
}
