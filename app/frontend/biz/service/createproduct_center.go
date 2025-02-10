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

type CreateproductCenterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateproductCenterService(Context context.Context, RequestContext *app.RequestContext) *CreateproductCenterService {
	return &CreateproductCenterService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateproductCenterService) Run(req *product_center.CreateproductCenterReq) (resp map[string]any, err error) {
	log.Printf("来到CreateproductCenterService函数这里了:")
	token := string(frontendUtils.GetTokenFromCtx(h.Context))
	userId := uint32(frontendUtils.GetUserIdFromCtx(h.Context))

	log.Printf("前端读取到的token:", token)
	log.Printf("前端读取到的userId:", userId)

	log.Printf("前端提交的表单数据:", req)

	productResp, err := rpc.ProductClient.CreateProduct(h.Context, &product.CreateProductReq{
		Token:           token,
		Name:            req.Name,
		Description:     req.Description,
		Picture:         req.Picture,
		Price:           req.Price,
		Categories:      req.Categories,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
		UserId:          int32(userId),
	})
	log.Printf("来到CreateproductCenterService函数成功调用后端，返回前端渲染:", productResp)

	// return utils.H{
	// 	"title": "product-center",
	// 	"items": productResp.Id,
	// }, nil
	resp = utils.H{
		"title":         "CreateProduct",
		"CreateProduct": productResp,
	}
	return resp, nil
	// redirect = "/"    // redirect string
	// if req.Next != "" {
	// 	redirect = req.Next
	// }
	// return
}
