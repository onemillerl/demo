package service

import (
	"context"
	"testing"

	user "gomall_demo/rpc_gen/kitex_gen/user"
)

func TestGetUserInfo_Run(t *testing.T) {
	ctx := context.Background()

	initUserClient()
	resp, err := userClient.GetUserInfo(ctx, &user.GetUserInfoReq{
		UserId: 4,
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
