package products

import (
	"context"
)

type ProductsInputPort interface {
	CreateProduct(ctx context.Context, request ProductRequest) (*ProductResponse, error)
}
