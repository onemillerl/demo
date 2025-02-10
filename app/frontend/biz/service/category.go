package service

import (
	"context"

	category "gomall_demo/app/frontend/hertz_gen/frontend/category"
	"gomall_demo/app/frontend/infra/rpc"
	"gomall_demo/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	hlog.Info("Initializing CategoryService")
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	hlog.Infof("CategoryService.Run called with request: %+v", req)

	p, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{CategoryName: req.Category})
	if err != nil {
		hlog.Errorf("Error fetching products for category '%s': %v", req.Category, err)

		return nil, err
	}
	// Prepare and return the response
	resp = utils.H{
		"title": "Category",
		"items": p.Products,
	}
	hlog.Infof("Response prepared: %+v", resp)
	return resp, nil
}
