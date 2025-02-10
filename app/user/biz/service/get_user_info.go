package service

import (
	"context"
	"log"

	"gomall_demo/app/user/biz/dal/mysql"
	"gomall_demo/app/user/biz/model"
	user "gomall_demo/rpc_gen/kitex_gen/user"
)

type GetUserInfoService struct {
	ctx context.Context
} // NewGetUserInfoService new GetUserInfoService
func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

// Run create note info
// Run retrieves the user information based on user ID.
func (s *GetUserInfoService) Run(req *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// 通过 UserId 获取用户信息
	log.Printf("来到后端的GetUserInfoService函数这里了:")
	log.Printf("要查询的用户id为:", req.UserId)

	userQuery := model.NewUserQuery(context.Background(), mysql.DB)
	userInfo, productid, err := userQuery.GetUserInfo(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	var productIds []int32
	for _, p := range userInfo.Products {
		productIds = append(productIds, p.Productid)
	}
	// 构建响应
	resp := &user.GetUserInfoResp{
		UserId:     int32(userInfo.ID), // 将数据库中的 ID 转换为 int32
		Email:      userInfo.Email,     // 返回用户的 Email
		Nickname:   userInfo.Nickname,
		Role:       userInfo.Role,
		ProductIds: productid,
	}
	log.Printf("完成后端的GetUserInfoService函数了:")

	return resp, nil
}
