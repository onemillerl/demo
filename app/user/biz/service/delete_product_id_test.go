package service

import (
	"context"
	"testing"
	user "gomall_demo/rpc_gen/kitex_gen/user"
)

func TestDeleteProductId_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteProductIdService(ctx)
	// init req and assert value

	req := &user.DeleteProductIdReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
