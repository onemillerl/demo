package category

import (
	"context"

	"gomall_demo/app/frontend/biz/service"
	category "gomall_demo/app/frontend/hertz_gen/frontend/category"

	"gomall_demo/app/frontend/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Category .
// @router /category/:category [GET]
func Category(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.CategoryReq
	hlog.Infof("Received category request: %s", c.Request.URI().String())

	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	hlog.Infof("Successfully processed category request. Response: %+v", resp)
	c.HTML(consts.StatusOK, "category", utils.WarpResponse(ctx, c, resp))
	hlog.Infof("Response sent successfully for category: %s", req.Category)
}
