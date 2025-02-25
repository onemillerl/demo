package service

import (
	"context"
	"gomall_demo/rpc_gen/kitex_gen/cart"
	order "gomall_demo/rpc_gen/kitex_gen/order"
	"testing"
)

func TestPlaceOrder_Run(t *testing.T) {
	ctx := context.Background()

	initOrderClient()
	resp, err := orderClient.PlaceOrder(ctx, &order.PlaceOrderReq{
		UserId:       3,
		UserCurrency: "good",
		Address: &order.Address{
			StreetAddress: "123 Main St",
			City:          "Anytown",
			State:         "CA",
			Country:       "China",
			ZipCode:       12345,
		},
		Email: "1234@qq.com",
		OrderItems: []*order.OrderItem{
			{
				Item: &cart.CartItem{
					ProductId: 1,
					Quantity:  1,
				},
				Cost: 100,
			},
		},
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
