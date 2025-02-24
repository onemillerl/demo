package service

import (
	userutils "gomall_demo/app/user/utils"
	"gomall_demo/rpc_gen/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var userClient userservice.Client

func initUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	userutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	userClient, err = userservice.NewClient("user", opts...)
	userutils.MustHandleError(err)
}
