package service

import (
	"context"
	"testing"

	user "gomall_demo/rpc_gen/kitex_gen/user"
)

func TestLogin_Run(t *testing.T) {
	ctx := context.Background()

	initUserClient()
	resp, err := userClient.Login(ctx, &user.LoginReq{
		Email:    "test20@126.com",
		Password: "123",
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
