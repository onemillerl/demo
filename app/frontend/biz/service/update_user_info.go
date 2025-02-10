package service

import (
	"context"
	"log"

	auth "gomall_demo/app/frontend/hertz_gen/frontend/auth"
	"gomall_demo/app/frontend/infra/rpc"
	frontendutils "gomall_demo/app/frontend/utils"
	gomallauth "gomall_demo/rpc_gen/kitex_gen/auth"
	"gomall_demo/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
)

type UpdateUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUserInfoService(Context context.Context, RequestContext *app.RequestContext) *UpdateUserInfoService {
	return &UpdateUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateUserInfoService) Run(req *auth.UpdateUserInfoReq) (resp map[string]any, err error) {
	// 通过 RPC 获取用户信息
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	log.Printf("确认密码PasswordConfirm: %s", req.PasswordConfirm)
	log.Printf("Email: %s", req.Email)

	respuser, err := rpc.UserClient.UpdateUserInfo(h.Context, &user.UpdateUserInfoReq{
		UserId:          int32(userId),
		Password:        req.Password,
		Email:           req.Email,
		PasswordConfirm: req.PasswordConfirm,
		Nickname:        req.Nickname,
		Role:            req.Role,
	})
	if err != nil {
		return nil, err
	}

	// 处理 RPC 可能返回空数据的情况
	if respuser == nil {
		return utils.H{
			"title":    "User Info",
			"userinfo": nil,
		}, nil
	}
	hlog.CtxInfof(h.Context, "getuserinfo response: %+v", resp)
	authtokenresp, err := rpc.AuthClient.DeliverTokenByRPC(h.Context, &gomallauth.DeliverTokenReq{
		UserId: respuser.UserId,
		Role:   respuser.Role,
	})
	// 存储用户信息到 session
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", respuser.UserId)
	session.Set("Role", respuser.Role)
	session.Set("token", authtokenresp.Token)
	err = session.Save()
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	userinfo := map[string]any{
		"UserId":   respuser.UserId,
		"Email":    respuser.Email,
		"Nickname": respuser.Nickname,
		"Role":     respuser.Role,
	}
	hlog.CtxInfof(h.Context, " userinfo: %+v", userinfo)
	log.Printf("Email: %+v", userinfo)

	// 返回给前端用于渲染页面
	return utils.H{
		"title":       "User Info",
		"getuserinfo": userinfo,
	}, nil
}
