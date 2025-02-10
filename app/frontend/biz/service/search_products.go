package service

import (
	"context"
	"fmt"

	product "gomall_demo/app/frontend/hertz_gen/frontend/product"
	rpcproduct "gomall_demo/rpc_gen/kitex_gen/product"

	"gomall_demo/app/frontend/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	fmt.Println("从前端传过来的 req.Q:", req.Q)
	products, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"items": products.Results,
		"q":     req.Q,
	}, nil
}
