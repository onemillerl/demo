package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "gomall_demo/app/frontend/hertz_gen/frontend/common"
	product_center "gomall_demo/app/frontend/hertz_gen/frontend/product_center"
)

type SearchproductCentersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchproductCentersService(Context context.Context, RequestContext *app.RequestContext) *SearchproductCentersService {
	return &SearchproductCentersService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchproductCentersService) Run(req *product_center.SearchproductCentersReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
