package service

import (
	"context"
	"testing"

	user "gomall_demo/rpc_gen/kitex_gen/user"
)

func TestRegister_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "test8@126.com", // 假设这是一个有效的邮箱
		Password:        "123",           // 假设这是正确的密码
		PasswordConfirm: "123",
		Nickname:        "demo",
		Role:            "admin",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
