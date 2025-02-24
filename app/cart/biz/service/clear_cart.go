package service

import (
	"context"
	"log"

	"gomall_demo/app/cart/biz/dal/mysql"
	"gomall_demo/app/cart/biz/model"
	cart "gomall_demo/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ClearCartService struct {
	ctx context.Context
} // NewClearCartService new ClearCartService
func NewClearCartService(ctx context.Context) *ClearCartService {
	return &ClearCartService{ctx: ctx}
}

// Run create note info
func (s *ClearCartService) Run(req *cart.ClearCartReq) (resp *cart.ClearCartResp, err error) {
	log.Printf("来到后端ClearCartService这里了:")
	log.Printf("要clear的用户ID: %d", req.UserId)

	err = model.EmptyCart(mysql.DB, s.ctx, req.GetUserId())
	if err != nil {
		return &cart.ClearCartResp{}, kerrors.NewBizStatusError(50001, "empty cart error")
	}

	return &cart.ClearCartResp{UserId: req.UserId}, nil
}
