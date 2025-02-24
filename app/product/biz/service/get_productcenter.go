package service

import (
	"context"
	"log"

	"gomall_demo/app/product/biz/dal/mysql"
	"gomall_demo/app/product/biz/model"
	"gomall_demo/app/product/infra/rpc"
	product "gomall_demo/rpc_gen/kitex_gen/product"
	"gomall_demo/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductcenterService struct {
	ctx context.Context
} // NewGetProductcenterService new GetProductcenterService
func NewGetProductcenterService(ctx context.Context) *GetProductcenterService {
	return &GetProductcenterService{ctx: ctx}
}

// Run create note info
func (s *GetProductcenterService) Run(req *product.GetProductcenterReq) (resp *product.GetProductcenterResp, err error) {
	log.Printf("来到后端的GetProductcenterService函数这里了:")

	if req.UserId == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "UserId is required")
	}

	userinforesp, err := rpc.UserClient.GetUserInfo(s.ctx, &user.GetUserInfoReq{
		UserId: req.UserId,
		// Role:   row.Role,
	})
	log.Printf("调用后端的GetUserInfo查询用户产品id: %v", userinforesp)

	if err != nil {
		return nil, err
	}
	productIds := userinforesp.ProductIds // 从 userinforesp 中获取产品 ID
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)

	var products []*product.Product
	// 遍历 ProductIds，依次查询商品
	for _, productId := range productIds {
		// Query the product from model.Product
		p, err := productQuery.GetById(int(productId))
		if err != nil {
			log.Printf("Error fetching product ID %d: %v", productId, err)
			continue // Log the error and continue querying other products
		}
		log.Printf("查找的商品id: %d", p.ID)
		if p.IsDeleted {
			return nil, kerrors.NewGRPCBizStatusError(2004006, "product has been deleted")
		}
		product := &product.Product{
			Id:          uint32(p.ID), // Assuming model.Product has field ID
			Name:        p.Name,       // Assuming model.Product has field Name
			Price:       p.Price,      // Assuming model.Product has field Price
			Description: p.Description,
			Picture:     p.Picture,
		}
		products = append(products, product)
	}
	log.Printf("后端的GetProductcenterService调用成功: %v", userinforesp)

	return &product.GetProductcenterResp{
		Results: products,
	}, nil
}
