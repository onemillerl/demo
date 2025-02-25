package service

import (
	"context"
	order "gomall_demo/rpc_gen/kitex_gen/order"
	"testing"
)

func TestListOrder_Run(t *testing.T) {
	ctx := context.Background()

	initOrderClient()
	resp, err := orderClient.ListOrder(ctx, &order.ListOrderReq{
		UserId: 3,
	})

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
