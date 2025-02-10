package utils

import (
	"context"

	"gomall_demo/app/frontend/infra/rpc"
	frontendutils "gomall_demo/app/frontend/utils"

	"gomall_demo/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/hertz/pkg/app"
	// "github.com/gorilla/sessions"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

// func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
// userId := frontendUtils.GetUserIdFromCtx(ctx)
// content["user_id"] = userId

// if userId > 0 {
// 	cartResp, err := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{
// 		UserId: uint32(userId),
// 	})
// 	if err == nil && cartResp != nil {
// 		content["cart_num"] = len(cartResp.Items)
// 	}
// }

// return content
// }

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	var cartNum int
	userId := frontendutils.GetUserIdFromCtx(ctx)
	cartResp, _ := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{UserId: userId})
	if cartResp != nil && cartResp.Cart != nil {
		cartNum = len(cartResp.Cart.Items)
	}
	content["user_id"] = ctx.Value(frontendutils.UserIdKey)
	content["cart_num"] = cartNum
	return content
}
