package service

import (
	"gomall_demo/rpc_gen/kitex_gen/checkout/checkoutservice"

	checkoututils "gomall_demo/app/checkout/utils"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var checkoutClient checkoutservice.Client

func initCheckoutClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	checkoututils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	checkoutClient, err = checkoutservice.NewClient("checkout", opts...)
	checkoututils.MustHandleError(err)
}
