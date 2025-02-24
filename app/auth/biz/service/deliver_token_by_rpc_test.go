package service

import (
	"context"
	"testing"

	auth "gomall_demo/rpc_gen/kitex_gen/auth"
)

func TestDeliverTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()

	initAuthClient()
	resp, err := authClient.DeliverTokenByRPC(ctx, &auth.DeliverTokenReq{
		UserId: 1,
		Role:   "admin",
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
