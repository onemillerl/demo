package service

import (
	"context"
	"errors"
	"log"
	"strconv"

	common "gomall_demo/app/frontend/hertz_gen/frontend/common"
	"gomall_demo/app/frontend/infra/rpc"
	frontendUtils "gomall_demo/app/frontend/utils"
	"gomall_demo/rpc_gen/kitex_gen/auth"
	"gomall_demo/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetproductCenterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetproductCenterService(Context context.Context, RequestContext *app.RequestContext) *GetproductCenterService {
	return &GetproductCenterService{RequestContext: RequestContext, Context: Context}
}

func (h *GetproductCenterService) Run(req *common.Empty) (resp map[string]any, err error) {
	userId := uint32(frontendUtils.GetUserIdFromCtx(h.Context))
	token := string(frontendUtils.GetTokenFromCtx(h.Context))
	log.Printf("来到run函数这里了:")

	log.Printf("product_Center的 userId:", userId)
	log.Printf("product_Center的token:", token)
	authtokenresp, err := rpc.AuthClient.VerifyTokenByRPC(h.Context, &auth.VerifyTokenReq{
		Token: token,
	})
	if err != nil {
		hlog.CtxErrorf(h.Context, "Failed to verify token: %v", err)
		return nil, err
	}
	if authtokenresp.Role != "seller" {
		hlog.CtxErrorf(h.Context, "Failed to enter Product-Center: %v, because your aren't seller!", authtokenresp)
		return nil, errors.New("unauthorized: user is not a seller")
	}
	productResp, err := rpc.ProductClient.GetProductcenter(h.Context, &product.GetProductcenterReq{
		UserId: int32(userId),
		Token:  token,
	})
	if err != nil {
		return nil, err
	}

	var items []map[string]string
	for _, v := range productResp.Results {
		item := map[string]string{
			"Name":        v.Name,
			"Description": v.Description,
			"Picture":     v.Picture,
			"Price":       strconv.FormatFloat(float64(v.Price), 'f', 2, 64),
			"Id":          strconv.FormatUint(uint64(v.Id), 10),
		}
		items = append(items, item)
	}
	return utils.H{
		"title": "Product-center",
		"items": items,
	}, nil
}
