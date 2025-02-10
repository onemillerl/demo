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
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// Finish your business logic.
	fmt.Println("从前端传过来的 Query:", req)

	// 基本校验
	if req.Name == "" || req.Description == "" || req.Picture == "" || req.Price < 0 || len(req.Categories) == 0 {
		return nil, errors.New("Name, Description, Picture ,Price or Categories is empty")
	}
	// 基本校验
	if req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("Password or PasswordConfirm is empty")
	}
	// 角色校验，只允许特定的角色
	authtokenresp, err := rpc.AuthClient.VerifyTokenByRPC(s.ctx, &auth.VerifyTokenReq{
		Token: req.Token,
		// Role:   row.Role,
	})
	if authtokenresp.Role != "seller" {
		hlog.CtxErrorf(s.ctx, "Failed to create Product: %v, because your aren't seller!", authtokenresp)
		return nil, err
	}

	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	newProduct, err := productQuery.CreateProduct(
		req.Name,        // 产品名称
		req.Description, // 产品描述
		req.Picture,     // 产品图片
		req.Price,       // 产品价格
		req.Categories,  // 产品类别（例如：["Electronics", "Laptop"]）
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %v", err)
	}
	log.Printf("生成的产品信息:", newProduct)
	updatedUserResp, err := rpc.UserClient.UpdateUserInfo(s.ctx, &user.UpdateUserInfoReq{
		UserId:          req.UserId,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
		Productid:       int32(newProduct.ID),
	})
	return &product.CreateProductResp{Id: uint32(newProduct.ID), Userid: updatedUserResp.UserId}, nil
}
