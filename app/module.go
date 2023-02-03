package app

import (
	"product-service-graphql/app/product"
	"product-service-graphql/app/product/usecase"
	repository "product-service-graphql/infrastructure/db/product"
	connection "product-service-graphql/internal/database"
)

type Module struct {
	productUseCase product.ProductUsecase
}

var (
	module *Module
)

func init() {
	repoProd := repository.NewProductDataHandler(connection.ProductDb)
	useProd := usecase.NewProductUsecase(repoProd)

	module = &Module{
		productUseCase: useProd,
	}
}

func GetProduct() *Module {
	if nil != module {
		return module
	}

	return nil
}
