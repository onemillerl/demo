package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"gomall_demo/app/product/biz/dal/mysql"
	"gomall_demo/app/product/biz/model"
	"gomall_demo/app/product/infra/rpc"
	"gomall_demo/rpc_gen/kitex_gen/auth"
	product "gomall_demo/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	log.Printf("更新商品，来到后端的UpdateProductService函数这里了:")
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}
	if req.Name == "" || req.Description == "" || req.Picture == "" || req.Price < 0 || len(req.Categories) == 0 {
		return nil, errors.New("Name, Description, Picture ,Price or Categories is empty")
	}
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	p, err := productQuery.GetById(int(req.Id))
	if err != nil {
		return nil, err
	}
	if p.IsDeleted {
		return nil, kerrors.NewGRPCBizStatusError(2004002, "product has been deleted")
	}
	authtokenresp, err := rpc.AuthClient.VerifyTokenByRPC(s.ctx, &auth.VerifyTokenReq{
		Token: req.Token,
	})
	if authtokenresp.Role != "seller" {
		hlog.CtxErrorf(s.ctx, "Failed to update Product: %v, because your aren't seller!", authtokenresp)
		return nil, err
	}
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	categories, err := categoryQuery.GetCategoriesByNames(req.Categories)
	if err != nil {
		return nil, fmt.Errorf("failed to find categories: %v", err)
	}

	// 更新商品信息
	updatedProduct := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  categories, // 关联新的 Categories
	}
	// 更新商品（在数据库中）
	if err := productQuery.UpdateProduct(int(req.Id), updatedProduct); err != nil {
		return nil, fmt.Errorf("failed to update product: %v", err)
	}
	// 返回更新后的商品信息
	return &product.UpdateProductResp{
		Id:   req.Id,
		Name: req.Name,
	}, nil
}
