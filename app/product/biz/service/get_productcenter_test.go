package service

import (
	"context"
	product "gomall_demo/rpc_gen/kitex_gen/product"
	"testing"
)

func TestGetProductcenter_Run(t *testing.T) {
	ctx := context.Background()

	initProductClient()
	resp, err := productClient.GetProductcenter(ctx, &product.GetProductcenterReq{
		UserId: 1,
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDAyOTY5OTMsInJvbGUiOiJhZG1pbiIsInVzZXJpZCI6MX0.mLoER8UiCuBPEsuDdDMzEb2Xe3Dxng5VfVtDulhblFQ",
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
