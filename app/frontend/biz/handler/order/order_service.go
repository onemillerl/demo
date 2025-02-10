package order

import (
	"context"

	"gomall_demo/app/frontend/biz/service"
	common "gomall_demo/app/frontend/hertz_gen/frontend/common"

	"gomall_demo/app/frontend/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// OrderList .
// @router /order [GET]
func OrderList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewOrderListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "order", utils.WarpResponse(ctx, c, resp))
}
