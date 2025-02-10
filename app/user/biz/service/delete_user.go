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

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	// 校验用户ID是否有效

	if req.UserId <= 0 || req.Password == "" {
		return nil, errors.New("UserId or password is empty")
	}

	row, err := model.GetByUserId(mysql.DB, s.ctx, req.UserId)
	log.Printf("GetByEmail row: %s,%d,%s", row.Email, row.ID, row.Role)

	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	// 在数据库中查找用户
	// row, err := model.GetByUserId(mysql.DB, s.ctx, req.UserId)
	// if err != nil {
	// 	return nil, err
	// }
	err = model.DeleteByUserId(mysql.DB, s.ctx, req.UserId)
	// 执行删除操作
	if err != nil {
		return nil, err
	}

	// 返回成功的响应
	resp = &user.DeleteUserResp{
		Success: true, // 成功删除用户
	}

	hlog.CtxInfof(s.ctx, "User with ID: %d successfully deleted", row.Nickname)
	return resp, nil
}
