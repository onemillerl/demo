package service

import (
	"context"
	product "gomall_demo/rpc_gen/kitex_gen/product"
)

type SearchproductcentersService struct {
	ctx context.Context
} // NewSearchproductcentersService new SearchproductcentersService
func NewSearchproductcentersService(ctx context.Context) *SearchproductcentersService {
	return &SearchproductcentersService{ctx: ctx}
}

// Run create note info
func (s *SearchproductcentersService) Run(req *product.SearchproductcentersReq) (resp *product.SearchproductcentersResp, err error) {
	// Finish your business logic.

	return
}
