package service

import (
	"context"
	product "gomall_demo/rpc_gen/kitex_gen/product"
	"testing"
)

func TestCreateProduct_Run(t *testing.T) {
	ctx := context.Background()

	initProductClient()
	resp, err := productClient.CreateProduct(ctx, &product.CreateProductReq{
		Name:            "New Product",
		Description:     "This is a new created product",
		Picture:         "/static/image/mouse-pad.jpeg",
		Price:           30,
		Categories:      []string{"sticker"},
		Token:           "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDAyOTkyNTUsInJvbGUiOiJzZWxsZXIiLCJ1c2VyaWQiOjZ9.rXyBbjCUEJUHjwNRJtpMi60v3L-P67UzuLY6LU4wM6A",
		UserId:          6,
		Password:        "123",
		PasswordConfirm: "123",
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
