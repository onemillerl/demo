package service

import (
	"context"
	"errors"
	"log"
	"strings"

	"gomall_demo/app/user/biz/dal/mysql"
	"gomall_demo/app/user/biz/model"
	"gomall_demo/app/user/infra/rpc"
	"gomall_demo/rpc_gen/kitex_gen/auth"
	user "gomall_demo/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info   （！PasswordConfirm）
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	hlog.CtxInfof(s.ctx, "Received register request: Email=%s, Nickname=%s, Role=%s", req.Email, req.Nickname, req.Role)

	// 基本校验
	if req.Email == "" || req.Password == "" || req.PasswordConfirm == "" || req.Nickname == "" || req.Role == "" {
		return nil, errors.New("email, password, nickname or role is empty")
	}
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("passwords do not match")
	}

	// 角色校验，只允许特定的角色
	validRoles := map[string]bool{"admin": true, "customer": true, "seller": true}
	if !validRoles[strings.ToLower(req.Role)] {
		return nil, errors.New("invalid role, must be 'admin', 'customer' or 'seller'")
	}

	// 加密密码
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		hlog.CtxErrorf(s.ctx, "Failed to hash password: %v", err)
		return nil, err
	}

	// 创建用户对象

	userQuery := model.NewUserQuery(context.Background(), mysql.DB)
	newUser, err := userQuery.CreateUser(req.Email, string(passwordHashed), req.Nickname, strings.ToLower(req.Role), []int32{})
	if err != nil {
		hlog.CtxErrorf(s.ctx, "Failed to create user: %v", err)
		return nil, err
	}

	log.Printf("User created successfully with ID: %d, Email: %s, Role: %s", newUser.ID, newUser.Email, newUser.Role)
	authresptoken, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &auth.DeliverTokenReq{
		UserId: int32(newUser.ID), // 使用已验证的用户 ID
		Role:   newUser.Role,
	})
	log.Printf("token: %s", authresptoken.Token)

	if authresptoken == nil || authresptoken.Token == "" {
		return nil, errors.New("failed to get token")
	}
	log.Printf("返回的注册resp: %+v", user.RegisterResp{UserId: int32(newUser.ID), Token: authresptoken.Token})

	return &user.RegisterResp{UserId: int32(newUser.ID), Token: authresptoken.Token}, nil
}
