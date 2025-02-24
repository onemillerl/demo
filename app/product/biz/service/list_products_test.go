package service

import (
	"context"
	product "gomall_demo/rpc_gen/kitex_gen/product"
	"testing"
)

func TestListProducts_Run(t *testing.T) {
	ctx := context.Background()

	initProductClient()
	resp, err := productClient.ListProducts(ctx, &product.ListProductsReq{
		CategoryName: "sticker",
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
