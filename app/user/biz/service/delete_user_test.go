package service

import (
	"context"
	user "gomall_demo/rpc_gen/kitex_gen/user"
	"testing"
)

func TestDeleteUser_Run(t *testing.T) {
	ctx := context.Background()

	initUserClient()
	resp, err := userClient.DeleteUser(ctx, &user.DeleteUserReq{
		UserId:   4,
		Password: "123",
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
