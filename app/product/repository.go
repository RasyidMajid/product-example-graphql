package product

import (
	"github.com/opentracing/opentracing-go"
	"product-service-graphql/app/product/model"
)

type ProductRepository interface {
	CreateProduct(span opentracing.Span, in model.Product) (model.ProductResponse, error)
	GetAllProduct(span opentracing.Span) (model.ProductResponse, error)
	GetProductByID(span opentracing.Span, id string) (model.ProductResponse, error)
}
