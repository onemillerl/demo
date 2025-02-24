package service

import (
	"context"
	product "gomall_demo/rpc_gen/kitex_gen/product"
	"testing"
)

func TestUpdateProduct_Run(t *testing.T) {
	ctx := context.Background()

	initProductClient()
	resp, err := productClient.UpdateProduct(ctx, &product.UpdateProductReq{
		Id:          2,
		Name:        "Mouse-Padddddddddd",
		Description: "The cloudwego mouse pad is a premium-grade accessory designed to enhance your computer usage experience. ",
		Picture:     "/static/image/mouse-pad.jpeg",
		Price:       188,
		Categories:  []string{"sticker"},
		Token:       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDAyOTkyNTUsInJvbGUiOiJzZWxsZXIiLCJ1c2VyaWQiOjZ9.rXyBbjCUEJUHjwNRJtpMi60v3L-P67UzuLY6LU4wM6A",
		UserId:      6,
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
