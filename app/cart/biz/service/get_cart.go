package service

import (
	"context"

	"gomall_demo/app/cart/biz/dal/mysql"
	"gomall_demo/app/cart/biz/model"
	cart "gomall_demo/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// resp = &cart.Cart{}
	// Finish your business logic.
	carts, err := model.GetCartByUserId(mysql.DB, s.ctx, req.GetUserId())
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	var items []*cart.CartItem
	for _, v := range carts {
		items = append(items, &cart.CartItem{ProductId: v.ProductId, Quantity: int32(v.Qty)})
	}

	return &cart.GetCartResp{Cart: &cart.Cart{UserId: req.GetUserId(), Items: items}}, nil
}
