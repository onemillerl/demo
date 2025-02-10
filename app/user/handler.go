package main

import (
	"context"

	"gomall_demo/app/user/biz/service"
	user "gomall_demo/rpc_gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoResp, err error) {
	resp, err = service.NewGetUserInfoService(ctx).Run(req)

	return resp, err
}

// UpdateUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq) (resp *user.UpdateUserInfoResp, err error) {
	resp, err = service.NewUpdateUserInfoService(ctx).Run(req)

	return resp, err
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserReq) (resp *user.DeleteUserResp, err error) {
	resp, err = service.NewDeleteUserService(ctx).Run(req)

	return resp, err
}

// DeleteProductId implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteProductId(ctx context.Context, req *user.DeleteProductIdReq) (resp *user.DeleteProductIdResp, err error) {
	resp, err = service.NewDeleteProductIdService(ctx).Run(req)

	return resp, err
}
