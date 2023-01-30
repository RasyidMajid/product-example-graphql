package di

import (
	"github.com/sarulabs/di/v2"
	"product-service-graphql/src/adapter/connection"
	repository "product-service-graphql/src/adapter/presenter/db/products"
	"product-service-graphql/src/usecase/products"
)

type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add([]di.Def{
		{Name: "product", Build: ProductUsecase},
	}...)
	return &Container{
		ctn: builder.Build(),
	}
}
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func ProductUsecase(_ di.Container) (interface{}, error) {
	repo := repository.NewProductDataHandler(connection.Product)
	return products.NewProductInteractor(repo), nil
}
