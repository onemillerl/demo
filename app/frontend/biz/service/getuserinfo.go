package service

import (
	"context"

	auth "gomall_demo/app/frontend/hertz_gen/frontend/auth"
	"gomall_demo/app/frontend/infra/rpc"
	frontendutils "gomall_demo/app/frontend/utils"
	"gomall_demo/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
)

type GetuserinfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetuserinfoService(Context context.Context, RequestContext *app.RequestContext) *GetuserinfoService {
	return &GetuserinfoService{RequestContext: RequestContext, Context: Context}
}

func (h *GetuserinfoService) Run(req *auth.GetuserinfoReq) (map[string]any, error) {
	// 通过 RPC 获取用户信息
	userId := frontendutils.GetUserIdFromCtx(h.Context)

	resp, err := rpc.UserClient.GetUserInfo(h.Context, &user.GetUserInfoReq{
		UserId: int32(userId),
	})
	if err != nil {
		return nil, err
	}

	// 处理 RPC 可能返回空数据的情况
	if resp == nil {
		return utils.H{
			"title":    "User Info",
			"userinfo": nil,
		}, nil
	}
	hlog.CtxInfof(h.Context, "getuserinfo response: %+v", resp)

	// 存储用户信息到 session
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	session.Set("Email", resp.Email)
	session.Set("Nickname", resp.Nickname)
	session.Set("Role", resp.Role)
	err = session.Save()
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	userinfo := map[string]any{
		"UserId":   resp.UserId,
		"Email":    resp.Email,
		"Nickname": resp.Nickname,
		"Role":     resp.Role,
	}
	hlog.CtxInfof(h.Context, " userinfo: %+v", userinfo)

	// 返回给前端用于渲染页面
	return utils.H{
		"title":       "User Info",
		"getuserinfo": userinfo,
	}, nil
}
