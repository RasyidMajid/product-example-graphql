package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"
	"os"
	cfg "product-service-graphql/config"
	"product-service-graphql/src/infrastructure/controller"
	"product-service-graphql/src/infrastructure/generated"
	container "product-service-graphql/src/shared/di"
	"product-service-graphql/src/shared/tracing"
	productUsecase "product-service-graphql/src/usecase/products"
)

func main() {

	config := cfg.GetConfig()

	fmt.Println(config)
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Server.Graphql.Port
	}

	ctn := container.NewContainer()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: Apply(ctn)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	tracer, closer := tracing.Init("Product-service")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func Apply(ctn *container.Container) *controller.Resolver {
	return &controller.Resolver{
		Product: ctn.Resolve("product").(*productUsecase.ProductInteractor),
	}
}
