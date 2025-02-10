package service

import (
	"context"
	"errors"
	"log"

	"gomall_demo/app/user/biz/dal/mysql"
	"gomall_demo/app/user/biz/model"
	user "gomall_demo/rpc_gen/kitex_gen/user"
)

type DeleteProductIdService struct {
	ctx context.Context
} // NewDeleteProductIdService new DeleteProductIdService
func NewDeleteProductIdService(ctx context.Context) *DeleteProductIdService {
	return &DeleteProductIdService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductIdService) Run(req *user.DeleteProductIdReq) (resp *user.DeleteProductIdResp, err error) {
	log.Printf("DeleteProductIdService要删除的productid: ", req.Id)

	if req.Id <= 0 {
		return nil, errors.New("id is empty")
	}
	ProductIdQuery := model.NewProductIdQuery(context.Background(), mysql.DB)
	err = ProductIdQuery.DeleteProductById(req.Id)
	if err != nil {
		return nil, errors.New("删除用户的商品Id失败")
	}
	log.Printf("成功DeleteProductIdService要删除的productid: ", req.Id)

	return &user.DeleteProductIdResp{
		Success: true, // 表示删除操作成功
	}, nil
}
