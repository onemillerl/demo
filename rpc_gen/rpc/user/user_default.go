package user

import (
	"context"
	user "gomall_demo/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetUserInfo(ctx context.Context, req *user.GetUserInfoReq, callOptions ...callopt.Option) (resp *user.GetUserInfoResp, err error) {
	resp, err = defaultClient.GetUserInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetUserInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq, callOptions ...callopt.Option) (resp *user.UpdateUserInfoResp, err error) {
	resp, err = defaultClient.UpdateUserInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateUserInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteUser(ctx context.Context, req *user.DeleteUserReq, callOptions ...callopt.Option) (resp *user.DeleteUserResp, err error) {
	resp, err = defaultClient.DeleteUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteProductId(ctx context.Context, req *user.DeleteProductIdReq, callOptions ...callopt.Option) (resp *user.DeleteProductIdResp, err error) {
	resp, err = defaultClient.DeleteProductId(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteProductId call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
