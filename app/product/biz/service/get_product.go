package service

import (
	"context"

	"gomall_demo/app/product/biz/dal/mysql"
	"gomall_demo/app/product/biz/model"
	product "gomall_demo/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)

	p, err := productQuery.GetById(int(req.Id))
	if err != nil {
		return nil, err
	}
	// 检查商品是否被删除
	if p.IsDeleted {
		return nil, kerrors.NewGRPCBizStatusError(2004006, "product has been deleted")
	}
	return &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Picture:     p.Picture,
			Price:       p.Price,
			Description: p.Description,
			Name:        p.Name,
		},
	}, nil
}
