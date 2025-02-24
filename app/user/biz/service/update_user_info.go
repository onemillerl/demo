package service

import (
	"context"
	"errors"
	"log"

	"gomall_demo/app/user/biz/dal/mysql"
	"gomall_demo/app/user/biz/model"
	user "gomall_demo/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserInfoService struct {
	ctx context.Context
} // NewUpdateUserInfoService new UpdateUserInfoService
func NewUpdateUserInfoService(ctx context.Context) *UpdateUserInfoService {
	return &UpdateUserInfoService{ctx: ctx}
}

// Run create note info
func (s *UpdateUserInfoService) Run(req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {
	// 校验用户ID是否有效
	log.Printf("准备调用 ,UserId: %d, Email: %s,Nickname: %s ,Role: %s,Password: %s, All req: %+v", req.UserId, req.Email, req.Nickname, req.Role, req.Password, req)
	log.Printf("确认密码PasswordConfirm: %s", req.PasswordConfirm)
	log.Printf("更新时候默认的productid: %d", req.Productid)

	if req.UserId <= 0 {
		err = errors.New("invalid user_id")
		hlog.CtxErrorf(s.ctx, "Invalid user_id: %d", req.UserId)
		return nil, err
	}
	// 基本校验
	if req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("password or PasswordConfirm is empty")
	}
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("passwords do not match")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	// 创建 UserQuery
	userQuery := model.NewUserQuery(context.Background(), mysql.DB)

	// 判断是否需要更新用户信息或添加商品
	if req.Productid == 0 {
		// 更新用户信息
		userinfo, err := userQuery.UpdateUser(mysql.DB, s.ctx, req.UserId, req.Email, string(passwordHashed), req.Nickname, req.Role, req.Productid)
		if err != nil {
			hlog.CtxErrorf(s.ctx, "Failed to update user with ID: %d, error: %v", req.UserId, err)
			return nil, err
		}
		// 返回更新后的用户信息
		resp = &user.UpdateUserInfoResp{
			UserId:   int32(userinfo.ID), // 返回更新后的用户数据
			Email:    userinfo.Email,
			Nickname: userinfo.Nickname,
			Role:     userinfo.Role,
		}
		return resp, nil
	} else {
		// 添加商品到用户
		userinfo, err := userQuery.AddProductToUser(mysql.DB, s.ctx, uint(req.UserId), req.Productid)
		if err != nil {
			hlog.CtxErrorf(s.ctx, "Failed to add product to user with ID: %d, error: %v", req.UserId, err)
			return nil, err
		}
		// 返回添加商品后的用户信息
		resp = &user.UpdateUserInfoResp{
			UserId:   int32(userinfo.ID), // 返回更新后的用户数据
			Email:    userinfo.Email,
			Nickname: userinfo.Nickname,
			Role:     userinfo.Role,
		}
		log.Printf("添加商品订单到用户成功的resp: %+v", resp)
		return resp, nil
	}
}
