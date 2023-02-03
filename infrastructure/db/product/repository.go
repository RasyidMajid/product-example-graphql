package product

import (
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"product-service-graphql/app/product/model"
)

type ProductDataHandler struct {
	db *gorm.DB
}

func NewProductDataHandler(db *gorm.DB) *ProductDataHandler {
	return &ProductDataHandler{
		db: db,
	}
}

func (d *ProductDataHandler) CreateProduct(span opentracing.Span, in model.Product) (model.ProductResponse, error) {
	panic(nil)
}
func (d *ProductDataHandler) GetAllProduct(span opentracing.Span) (model.ProductResponse, error) {
	panic(nil)
}
func (d *ProductDataHandler) GetProductByID(span opentracing.Span, id string) (model.ProductResponse, error) {
	panic(nil)
}
