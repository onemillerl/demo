package service

import (
	"context"
	cart "gomall_demo/rpc_gen/kitex_gen/cart"
	"testing"
)

func TestEmptyCart_Run(t *testing.T) {
	ctx := context.Background()

	initCartClient()
	resp, err := cartClient.EmptyCart(ctx, &cart.EmptyCartReq{
		UserId: 3,
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
