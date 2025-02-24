package service

import (
	"context"
	product "gomall_demo/rpc_gen/kitex_gen/product"
	"testing"
)

func TestGetProduct_Run(t *testing.T) {
	ctx := context.Background()

	initProductClient()
	resp, err := productClient.GetProduct(ctx, &product.GetProductReq{
		Id: 1,
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
