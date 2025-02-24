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
	"gomall_demo/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	log.Printf("删除商品，来到后端的DeleteProductService函数这里了:")
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}
	authtokenresp, err := rpc.AuthClient.VerifyTokenByRPC(s.ctx, &auth.VerifyTokenReq{
		Token: req.Token,
	})
	if err != nil {
		hlog.CtxErrorf(s.ctx, "Failed to verify token: %v", err)
		return nil, err
	}
	// 验证用户角色是否是 "seller"
	if authtokenresp.Role != "seller" {
		hlog.CtxErrorf(s.ctx, "Failed to delete Product: %v, because your aren't seller!", authtokenresp)
		return nil, errors.New("You don't have permission to delete products")
	}
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	checkproduct, err := productQuery.GetById(int(req.Id))
	if err != nil {
		return nil, fmt.Errorf("Failed to find product: %v", err)
	}
	err = productQuery.SoftDeleteProduct(int(req.Id))
	if err != nil {
		return nil, fmt.Errorf("Failed to soft delete product: %v", err)
	}
	// 更新user 表中关联的产品ID表

	userinforesp, err := rpc.UserClient.DeleteProductId(s.ctx, &user.DeleteProductIdReq{
		Id: int32(req.Id),
		// Role:   row.Role,
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to  更新user 表中关联的产品ID表t: %v", err)
	}
	log.Printf("返回删除的商品情况: %v", userinforesp)

	return &product.DeleteProductResp{
		Id:   uint32(req.Id),
		Name: checkproduct.Name,
	}, nil
}
