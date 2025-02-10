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

type DeleteproductCenterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteproductCenterService(Context context.Context, RequestContext *app.RequestContext) *DeleteproductCenterService {
	return &DeleteproductCenterService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteproductCenterService) Run(req *product_center.DeleteproductCenterReq) (resp map[string]any, err error) {
	log.Printf("删除UpdateproductCenterService,删除的产品id:", req.Id)
	token := string(frontendUtils.GetTokenFromCtx(h.Context))
	delproductResp, err := rpc.ProductClient.DeleteProduct(h.Context, &product.DeleteProductReq{
		Id:    req.Id,
		Token: token,
	})
	log.Printf("成功删除的产品id:", delproductResp.Id)
	resp = utils.H{
		"title":         "DeleteProduct",
		"DeleteProduct": delproductResp,
	}
	return resp, nil
}
