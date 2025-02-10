package service

import (
	"context"

	common "gomall_demo/app/frontend/hertz_gen/frontend/common"
	"gomall_demo/app/frontend/infra/rpc"
	"gomall_demo/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

// func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
// 	//defer func() {
// 	// hlog.CtxInfof(h.Context, "req = %+v", req)
// 	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
// 	//}()
// 	// todo edit your code
// 	resp := make(map[string]any)
// 	items := []map[string]any{
// 		{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/t-shirt-1.jpeg"},
// 		{"Name": "T-shirt-2", "Price": 160, "Picture": "/static/image/t-shirt-1.jpeg"},
// 		{"Name": "T-shirt-3", "Price": 140, "Picture": "/static/image/sweatshirt.jpeg"},
// 		{"Name": "T-shirt-4", "Price": 130, "Picture": "/static/image/t-shirt-1.jpeg"},
// 		{"Name": "T-shirt-5", "Price": 120, "Picture": "/static/image/t-shirt-1.jpeg"},
// 		{"Name": "T-shirt-6", "Price": 120, "Picture": "/static/image/t-shirt-1.jpeg"},
// 	}
// 	resp["Title"] = "Hot Sales"
// 	resp["Items"] = items
// 	return resp, nil
// }

func (h *HomeService) Run(req *common.Empty) (res map[string]any, err error) {
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": "Hot sale",
		"items": products.Products,
	}, nil
}
