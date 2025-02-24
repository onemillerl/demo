package service

import (
	authutils "gomall_demo/app/auth/utils"
	"gomall_demo/rpc_gen/kitex_gen/auth/authservice"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var authClient authservice.Client

func initAuthClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	authutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	authClient, err = authservice.NewClient("auth", opts...)
	authutils.MustHandleError(err)
}
