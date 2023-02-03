package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/plugin/federation/testdata/entityresolver/generated"
	"log"
	"net/http"
	"os"
	"product-service-graphql/infrastructure/gql/resolver"
	ConfigServer "product-service-graphql/internal/server"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = ConfigServer.ConfigServer.Graphql.Port
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: Apply(),
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
func Apply() *resolver.Resolver {
	return &resolver.Resolver{}
}
