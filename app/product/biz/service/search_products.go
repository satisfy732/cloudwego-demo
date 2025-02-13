package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/product/biz/model"
	product "github.com/cloudwego/biz-demo/gomall/app/product/kitex_gen/product"
	"github.com/cloudwego/biz-demo/gomall/app/product/biz/dal/mysql"
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
    productQuery := model.NewProductQuery(s.ctx,mysql.DB)
	products,err := productQuery.SearchProducts(req.Query)
	var results []*product.Product
	for _,v :=range products{
		results = append (results,&product.Product{
			Id: uint32(v.ID),
				Picture : v.Picture,
				Price: v.Price,
				Description: v.Description,
				Name: v.Name,
		})
	}
	return &product.SearchProductsResp{
		Results: results,
	},err
}
