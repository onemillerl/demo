package service

import (
	"context"
	"log"

	product_center "gomall_demo/app/frontend/hertz_gen/frontend/product_center"
	"gomall_demo/app/frontend/infra/rpc"
	frontendUtils "gomall_demo/app/frontend/utils"
	"gomall_demo/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateproductCenterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateproductCenterService(Context context.Context, RequestContext *app.RequestContext) *UpdateproductCenterService {
	return &UpdateproductCenterService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateproductCenterService) Run(req *product_center.UpdateproductCenterReq) (resp map[string]any, err error) {
	log.Printf("更新UpdateproductCenterService")
	log.Printf("更新UpdateproductCenterService,更新的产品id:", req.Id)
	log.Printf("前端有id")

	userId := uint32(frontendUtils.GetUserIdFromCtx(h.Context))
	token := string(frontendUtils.GetTokenFromCtx(h.Context))
	log.Printf("userid 和 token", userId, token)

	productResp, err := rpc.ProductClient.UpdateProduct(h.Context, &product.UpdateProductReq{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  req.Categories,
		Token:       token,
		UserId:      int32(userId),
	})
	log.Printf("成功更新的产品id:", productResp.Id)
	resp = utils.H{
		"title":         "UpdateProduct",
		"UpdateProduct": productResp,
	}
	return resp, nil
}
