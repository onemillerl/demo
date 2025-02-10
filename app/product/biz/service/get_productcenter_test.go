package service

import (
	"context"
	"testing"
	product "gomall_demo/rpc_gen/kitex_gen/product"
)

func TestGetProductcenter_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetProductcenterService(ctx)
	// init req and assert value

	req := &product.GetProductcenterReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
