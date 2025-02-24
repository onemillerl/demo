package service

import (
	productutils "gomall_demo/app/product/utils"
	"gomall_demo/rpc_gen/kitex_gen/product/productcatalogservice"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var productClient productcatalogservice.Client

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	productutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	productClient, err = productcatalogservice.NewClient("product", opts...)
	productutils.MustHandleError(err)
}
