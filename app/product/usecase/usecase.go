package usecase

import (
	"context"
	"product-service-graphql/app/product"
	"product-service-graphql/app/product/model"
)

type ProductUsecase struct {
	repository product.ProductRepository
}

func NewProductUsecase(r product.ProductRepository) product.ProductUsecase {
	return &ProductUsecase{
		repository: r,
	}
}

func (u *ProductUsecase) CreateProduct(ctx context.Context, in model.Product) (model.ProductResponse, error) {
	panic(nil)
}
func (u *ProductUsecase) GetAllProduct(ctx context.Context) (model.ProductResponse, error) {
	panic(nil)
}
func (u *ProductUsecase) GetProductByID(ctx context.Context, id string) (model.ProductResponse, error) {
	panic(nil)
}
