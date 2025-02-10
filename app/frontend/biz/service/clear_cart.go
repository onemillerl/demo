package service

import (
	"context"
	"log"

	common "gomall_demo/app/frontend/hertz_gen/frontend/common"
	"gomall_demo/app/frontend/infra/rpc"
	frontendUtils "gomall_demo/app/frontend/utils"
	"gomall_demo/rpc_gen/kitex_gen/cart"
	"gomall_demo/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type ClearCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewClearCartService(Context context.Context, RequestContext *app.RequestContext) *ClearCartService {
	return &ClearCartService{RequestContext: RequestContext, Context: Context}
}

func (h *ClearCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	log.Printf("准备发送请求到后端,来到ClearCartService这里了:")
	cleancartresp, err := rpc.CartClient.ClearCart(h.Context, &cart.ClearCartReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
	})
	userinforesp, err := rpc.UserClient.GetUserInfo(h.Context, &user.GetUserInfoReq{
		UserId: int32(cleancartresp.UserId),
	})
	resp = utils.H{
		"title":     "CleanCart",
		"CleanCart": userinforesp,
	}
	return resp, nil
}
