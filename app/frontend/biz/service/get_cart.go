package service

import (
	"context"
	"strconv"

	frontendUtils "gomall_demo/app/frontend/utils"

	common "gomall_demo/app/frontend/hertz_gen/frontend/common"
	"gomall_demo/app/frontend/infra/rpc"
	"gomall_demo/rpc_gen/kitex_gen/cart"
	"gomall_demo/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	carts, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
	})
	if err != nil {
		return nil, err
	}

	var items []map[string]string
	var total float32
	for _, v := range carts.Cart.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: v.ProductId})
		if err != nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{"Name": p.Name, "Description": p.Description, "Picture": p.Picture, "Price": strconv.FormatFloat(float64(p.Price), 'f', 2, 64), "Qty": strconv.Itoa(int(v.Quantity))})
		total += float32(v.Quantity) * p.Price
	}
	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
