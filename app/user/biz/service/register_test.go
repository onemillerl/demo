package service

import (
	"context"
	"testing"

	user "gomall_demo/rpc_gen/kitex_gen/user"
)

func TestRegister_Run(t *testing.T) {
	ctx := context.Background()

	initUserClient()
	resp, err := userClient.Register(ctx, &user.RegisterReq{
		Email:           "cccccf@qq.com", // 假设这是一个有效的邮箱
		Password:        "123",           // 假设这是正确的密码
		PasswordConfirm: "123",
		Nickname:        "ccccf",
		Role:            "seller",
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
