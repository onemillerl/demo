package service

import (
	"context"
	"log"

	auth "gomall_demo/app/frontend/hertz_gen/frontend/auth"
	common "gomall_demo/app/frontend/hertz_gen/frontend/common"
	"gomall_demo/app/frontend/infra/rpc"
	frontendutils "gomall_demo/app/frontend/utils"
	"gomall_demo/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
)

type DeleteUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteUserService(Context context.Context, RequestContext *app.RequestContext) *DeleteUserService {
	return &DeleteUserService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteUserService) Run(req *auth.DeleteUserReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// 通过 RPC 获取用户信息
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	log.Printf("要删除的用户id: %d", userId)

	result, err := rpc.UserClient.DeleteUser(h.Context, &user.DeleteUserReq{
		UserId:   int32(userId),
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	if result.Success == true {
		hlog.CtxInfof(h.Context, "注销成功: %+v", result)
	}

	session := sessions.Default(h.RequestContext)
	session.Clear()
	err = session.Save()
	if err != nil {
		return nil, err
	}

	// 返回给前端用于渲染页面
	return
}
