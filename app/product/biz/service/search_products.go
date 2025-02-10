package service

import (
	"context"
	"fmt"
	"log"

	"gomall_demo/app/product/biz/dal/mysql"
	"gomall_demo/app/product/biz/model"
	product "gomall_demo/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.

	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	log.Printf("productQuery : %+v", productQuery)
	log.Printf("从前端传过来的Query : %+s", req.Query)
	fmt.Println("productQuery:", productQuery)
	fmt.Println("从前端传过来的 Query:", req.Query)

	products, err := productQuery.SearchProducts(req.Query)
	var results []*product.Product
	for _, v := range products {
		results = append(results, &product.Product{
			Id:          uint32(v.ID),
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}

	return &product.SearchProductsResp{Results: results}, err
}
