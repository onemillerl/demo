package product_center

import (
	"context"
	"log"

	"gomall_demo/app/frontend/biz/service"
	"gomall_demo/app/frontend/biz/utils"
	common "gomall_demo/app/frontend/hertz_gen/frontend/common"
	product_center "gomall_demo/app/frontend/hertz_gen/frontend/product_center"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetproductCenter .
// @router /product_center [GET]
func GetproductCenter(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetproductCenterService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "productcenter", utils.WarpResponse(ctx, c, resp))
}

// SearchproductCenters .
// @router /search [GET]
func SearchproductCenters(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product_center.SearchproductCentersReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &common.Empty{}
	resp, err = service.NewSearchproductCentersService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CreateproductCenter .
// @router /product_center [POST]
func CreateproductCenter(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product_center.CreateproductCenterReq

	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	resp, err := service.NewCreateproductCenterService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	// c.HTML(consts.StatusOK, "productcenter", resp)
	c.HTML(consts.StatusOK, "createproduct", utils.WarpResponse(ctx, c, resp))

	// c.Redirect(consts.StatusOK, []byte(resp)) // 重定向
}

// UpdateproductCenter .
// @router /product_center [PUT]
func UpdateproductCenter(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product_center.UpdateproductCenterReq
	err = c.BindAndValidate(&req)
	log.Printf("来的handler这里")

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	resp, err := service.NewUpdateproductCenterService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "updateproduct", utils.WarpResponse(ctx, c, resp))

	// c.Redirect(consts.StatusOK, []byte(resp)) // 重定向
}

// DeleteproductCenter .
// @router /product_center [DELETE]
func DeleteproductCenter(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product_center.DeleteproductCenterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewDeleteproductCenterService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "deleteproduct", utils.WarpResponse(ctx, c, resp))
}
