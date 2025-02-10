package service

import (
	"context"
	"time"

	"gomall_demo/app/frontend/infra/rpc"
	"gomall_demo/app/frontend/types"
	frontendutils "gomall_demo/app/frontend/utils"
	rpcorder "gomall_demo/rpc_gen/kitex_gen/order" // 还没生成
	rpcproduct "gomall_demo/rpc_gen/kitex_gen/product"

	common "gomall_demo/app/frontend/hertz_gen/frontend/common"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	var orders []*types.Order
	listOrderResp, err := rpc.OrderClient.ListOrder(h.Context, &rpcorder.ListOrderReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	if listOrderResp == nil || len(listOrderResp.Orders) == 0 {
		return utils.H{
			"title":  "Order",
			"orders": orders,
		}, nil
	}

	for _, v := range listOrderResp.Orders {
		var items []types.OrderItem
		var total float32
		if len(v.OrderItems) > 0 {
			for _, vv := range v.OrderItems {
				total += vv.Cost
				i := vv.Item
				productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: i.ProductId})
				if err != nil {
					return nil, err
				}
				if productResp.Product == nil {
					continue
				}
				p := productResp.Product
				items = append(items, types.OrderItem{
					ProductId:   i.ProductId,
					Qty:         uint32(i.Quantity),
					ProductName: p.Name,
					Picture:     p.Picture,
					Cost:        vv.Cost,
				})
			}
		}
		timeObj := time.Unix(int64(v.CreatedAt), 0)
		orders = append(orders, &types.Order{
			Cost:        total,
			Items:       items,
			CreatedDate: timeObj.Format("2006-01-02 15:04:05"),
			OrderId:     v.OrderId,
			Consignee:   types.Consignee{Email: v.Email},
		})
	}

	return utils.H{
		"title":  "Order",
		"orders": orders,
	}, nil
}
