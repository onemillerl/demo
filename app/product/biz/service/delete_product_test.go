package service

import (
	"context"
	product "gomall_demo/rpc_gen/kitex_gen/product"
	"testing"
)

func TestDeleteProduct_Run(t *testing.T) {
	ctx := context.Background()

	initProductClient()
	resp, err := productClient.DeleteProduct(ctx, &product.DeleteProductReq{
		Id:    8,
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDAyOTkyNTUsInJvbGUiOiJzZWxsZXIiLCJ1c2VyaWQiOjZ9.rXyBbjCUEJUHjwNRJtpMi60v3L-P67UzuLY6LU4wM6A",
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
