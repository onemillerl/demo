package rpc

import (
	"sync"

	"gomall_demo/app/product/conf"
	productutils "gomall_demo/app/product/utils"
	"gomall_demo/rpc_gen/kitex_gen/auth/authservice"

	"gomall_demo/rpc_gen/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	AuthClient authservice.Client
	UserClient userservice.Client
	once       sync.Once
)

func InitClient() {
	once.Do(func() {
		initAuthClient()
		initUserClient()
	})
}

func initAuthClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	productutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	AuthClient, err = authservice.NewClient("auth", opts...)
	productutils.MustHandleError(err)
}

func initUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	productutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	UserClient, err = userservice.NewClient("user", opts...)
	productutils.MustHandleError(err)
}
