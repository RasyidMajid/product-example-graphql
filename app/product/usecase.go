package product

import (
	"context"
	"product-service-graphql/app/product/model"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, in model.Product) (model.ProductResponse, error)
	GetAllProduct(ctx context.Context) (model.ProductResponse, error)
	GetProductByID(ctx context.Context, id string) (model.ProductResponse, error)
}
