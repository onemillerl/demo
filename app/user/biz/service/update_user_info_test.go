package service

import (
	"context"
	user "gomall_demo/rpc_gen/kitex_gen/user"
	"testing"
)

func TestUpdateUserInfo_Run(t *testing.T) {
	ctx := context.Background()

	initUserClient()
	resp, err := userClient.UpdateUserInfo(ctx, &user.UpdateUserInfoReq{
		UserId:          4,
		Email:           "test20@126.com",
		Password:        "123",
		PasswordConfirm: "123",
		Nickname:        "demo",
		Role:            "admin",
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
