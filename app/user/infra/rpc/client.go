package rpc

import (
	"sync"

	"gomall_demo/app/user/conf"
	userutils "gomall_demo/app/user/utils"
	"gomall_demo/rpc_gen/kitex_gen/auth/authservice"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	AuthClient authservice.Client
	once       sync.Once
)

func InitClient() {
	once.Do(func() {
		initAuthClient()
	})
}

func initAuthClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	userutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	AuthClient, err = authservice.NewClient("auth", opts...)
	userutils.MustHandleError(err)
}
