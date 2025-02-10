package service

import (
	"context"
	"log"

	auth "gomall_demo/app/frontend/hertz_gen/frontend/auth"

	common "gomall_demo/app/frontend/hertz_gen/frontend/common"

	"gomall_demo/app/frontend/infra/rpc"
	"gomall_demo/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sessions"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	res, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
		Nickname:        req.Nickname,
		Role:            req.Role,
	})
	hlog.CtxInfof(h.Context, "前端的数据Sending register request: Email=%s, Password=%s, ConfirmPassword=%s", req.Email, req.Password, req.PasswordConfirm)
	log.Printf("接受返回的注册resp:", res)
	// hlog.CtxErrorf(h.Context, "网页获取的上下文信息userRespUserId: %v", res.UserId, "网页获取的上下文信息userRespToken: %v", res.Token)
	if err != nil {
		// 打印详细错误信息
		hlog.CtxErrorf(h.Context, "Failed to register user. Error: %v", err)
		return nil, err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", res.UserId)
	session.Set("token", res.Token)
	err = session.Save()
	if err != nil {
		hlog.CtxErrorf(h.Context, "sessions error Failed to register user. Error: %v", err)
		return nil, err
	}
	return
}
