package rpc

import (
	"sync"

	"gomall_demo/app/cart/conf"
	cartutils "gomall_demo/app/cart/utils"
	"gomall_demo/rpc_gen/kitex_gen/product/productcatalogservice"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	UserClient    userservice.Client

	once sync.Once
)

func InitClient() {
	once.Do(func() {
		initProductClient()
		initUserClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}

func initUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	UserClient, err = userservice.NewClient("user", opts...)
	cartutils.MustHandleError(err)
}
