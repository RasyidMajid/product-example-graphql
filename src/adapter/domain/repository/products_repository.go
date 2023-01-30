package repository

import (
	"github.com/opentracing/opentracing-go"
	"product-service-graphql/src/adapter/domain/entity"
)

type ProductsRepository interface {
	CreateProduct(span opentracing.Span, request entity.ProductRequest) (*entity.ProductResponse, error)
}
