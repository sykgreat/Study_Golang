package service

import "context"

var ProductService = &productService{}

type productService struct {
}

func (p *productService) mustEmbedUnimplementedProductServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (p *productService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	stock := GetProductByID(request.ProductId)
	return &ProductResponse{ProductStock: stock}, nil
}

func GetProductByID(id int32) int32 {
	return id
}
