package service

import (
	"context"
	cart "gomall_demo/rpc_gen/kitex_gen/cart"
	"testing"
)

func TestAddItem_Run(t *testing.T) {
	ctx := context.Background()

	initCartClient()
	resp, err := cartClient.AddItem(ctx, &cart.AddItemReq{
		UserId: 3,
		Item: &cart.CartItem{
			ProductId: 1,
			Quantity:  1,
		},
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
