package products

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"product-service-graphql/src/adapter/domain/entity"
	"product-service-graphql/src/adapter/domain/repository"
	"product-service-graphql/src/shared/tracing"
)

type Product struct {
	ID           *string
	NamaProduct  *string
	JenisProduct *string
	HargaProduct *float64
}
type ProductRequest struct {
	ID           string
	NamaProduct  string
	JenisProduct string
	HargaProduct float64
}
type ProductResponse struct {
	StatusCode string
	StatusDesc string
	Data       *Product
	DataList   []*Product
}

type ProductInteractor struct {
	repo repository.ProductsRepository
}

func NewProductInteractor(r repository.ProductsRepository) *ProductInteractor {
	return &ProductInteractor{
		repo: r,
	}
}

func (u *ProductInteractor) CreateProduct(ctx context.Context, request ProductRequest) (*ProductResponse, error) {
	sp := tracing.CreateChildSpan(ctx, string(tracing.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, request)

	reqData := entity.ProductRequest{
		NamaProduct:  request.NamaProduct,
		JenisProduct: request.JenisProduct,
		HargaProduct: request.HargaProduct,
	}

	res, err := u.repo.CreateProduct(sp, reqData)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var out *ProductResponse
	err = mapstructure.Decode(res, &out)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, out)
	return out, nil
}
