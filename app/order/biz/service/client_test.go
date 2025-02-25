package service

import (
	"gomall_demo/rpc_gen/kitex_gen/order/orderservice"

	orderutils "gomall_demo/app/order/utils"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var orderClient orderservice.Client

func initOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	orderutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	orderClient, err = orderservice.NewClient("order", opts...)
	orderutils.MustHandleError(err)
}
