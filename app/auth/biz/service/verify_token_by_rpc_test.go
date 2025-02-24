package service

import (
	"context"
	auth "gomall_demo/rpc_gen/kitex_gen/auth"
	"testing"
)

func TestVerifyTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()

	initAuthClient()
	resp, err := authClient.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDAyOTY5OTMsInJvbGUiOiJhZG1pbiIsInVzZXJpZCI6MX0.mLoER8UiCuBPEsuDdDMzEb2Xe3Dxng5VfVtDulhblFQ",
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
