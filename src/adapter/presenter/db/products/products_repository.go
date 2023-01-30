package products

import (
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"product-service-graphql/src/adapter/domain/entity"
	"product-service-graphql/src/adapter/presenter/model"
	"product-service-graphql/src/shared/tracing"
)

type ProductDataHandler struct {
	db *gorm.DB
}

func NewProductDataHandler(db *gorm.DB) *ProductDataHandler {
	return &ProductDataHandler{
		db: db,
	}
}

func (d *ProductDataHandler) CreateProduct(span opentracing.Span, in entity.ProductRequest) (*entity.ProductResponse, error) {
	sp := tracing.CreateSubChildSpan(span, "CreateProduct")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	data := &model.ProductModel{
		NamaProduct:  in.NamaProduct,
		JenisProduct: in.JenisProduct,
		HargaProduct: in.HargaProduct,
	}

	err := d.db.Debug().Save(data).Error
	if err != nil {
		return nil, err
	}
	out := &entity.ProductResponse{
		StatusCode: "00",
		StatusDesc: "Success",
		Data:       nil,
		DataList:   nil,
	}

	tracing.LogResponse(sp, out)
	return out, nil
}
