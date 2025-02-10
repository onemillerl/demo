package service

import (
	"context"
	"testing"
	product "gomall_demo/rpc_gen/kitex_gen/product"
)

func TestSearchproductcenters_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchproductcentersService(ctx)
	// init req and assert value

	req := &product.SearchproductcentersReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
