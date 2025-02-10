package service

import (
	"context"
	"testing"
	cart "gomall_demo/rpc_gen/kitex_gen/cart"
)

func TestClearCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewClearCartService(ctx)
	// init req and assert value

	req := &cart.ClearCartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
