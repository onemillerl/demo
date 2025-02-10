package service

import (
	"context"
	"testing"

	user "gomall_demo/rpc_gen/kitex_gen/user"
)

func TestLogin_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	// 3. 模拟一个有效的请求
	req := &user.LoginReq{
		Email:    "test3@126.com", // 假设这是一个有效的邮箱
		Password: "123",           // 假设这是正确的密码
	}

	// 4. 调用 Run 方法执行登录逻辑
	resp, err := s.Run(req)

	// 5. 输出错误和响应信息（便于调试）
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// 6. 使用断言来验证响应的正确性

	// 验证是否没有错误
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// 验证响应中是否包含 user_id 和 Token
	if resp == nil {
		t.Errorf("Expected a valid response, but got nil")
	} else {
		if resp.UserId == 0 {
			t.Errorf("Expected a non-zero user_id, but got: %d", resp.UserId)
		}
		if resp.Token == "" {
			t.Errorf("Expected a token, but got empty string")
		}
	}
}
