package service

import (
	"context"

	cart "gomall_demo/app/frontend/hertz_gen/frontend/cart"
	common "gomall_demo/app/frontend/hertz_gen/frontend/common"
	"gomall_demo/app/frontend/infra/rpc"
	frontendUtils "gomall_demo/app/frontend/utils"
	rpccart "gomall_demo/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.ProductNum,
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
