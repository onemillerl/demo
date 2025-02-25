package service

import (
	"gomall_demo/rpc_gen/kitex_gen/payment/paymentservice"

	paymentutils "gomall_demo/app/payment/utils"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var paymentClient paymentservice.Client

func initPaymentClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver("172.30.137.28:8500")
	paymentutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	paymentClient, err = paymentservice.NewClient("payment", opts...)
	paymentutils.MustHandleError(err)
}
