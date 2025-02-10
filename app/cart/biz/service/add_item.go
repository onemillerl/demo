package service

import (
	"context"
	"log"

	"gomall_demo/app/cart/biz/dal/mysql"
	"gomall_demo/app/cart/biz/model"
	"gomall_demo/app/cart/infra/rpc"
	cart "gomall_demo/rpc_gen/kitex_gen/cart"
	product "gomall_demo/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	getProduct, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.GetProductId()})
	if err != nil {
		return nil, err
	}
	log.Printf("这里的产品rpc正确调用了getProduct id : %d", getProduct.Product.Id)

	if getProduct.Product == nil || getProduct.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not exist")
	}

	err = model.AddCart(mysql.DB, s.ctx, &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       uint32(req.Item.Quantity),
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}

	return &cart.AddItemResp{}, nil
}
