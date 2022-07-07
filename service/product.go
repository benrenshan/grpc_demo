package service

import "context"

var ProductService = &productService{}

type productService struct {
}

func (p *productService) mustEmbedUnimplementedProdServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (p *productService) GetProductStock(context context.Context, request *ProductRequest) (*ProductResponse, error) {
	//实现具体的业务逻辑
	//例如，查询当前的库存数
	stock := p.GetStockByID(request.GetProdId())
	return &ProductResponse{ProdStock: stock}, nil
}

func (p *productService) GetStockByID(id int32) int32 {
	//写一个假数据
	return 100
}
