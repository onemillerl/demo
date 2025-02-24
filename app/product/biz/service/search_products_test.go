package service

import (
	"context"
	product "gomall_demo/rpc_gen/kitex_gen/product"
	"testing"
)

func TestSearchProducts_Run(t *testing.T) {
	ctx := context.Background()

	initProductClient()
	resp, err := productClient.SearchProducts(ctx, &product.SearchProductsReq{
		Query: "t-shirt",
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
