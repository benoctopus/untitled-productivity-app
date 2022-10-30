package main

import (
	"gateway/config"
	"gateway/di"
	"gateway/graph"
	"gateway/graph/generated"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	di.Init()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(di.Invoke(func(conf *config.Config) error {
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", conf.Port)
		return http.ListenAndServe(":"+conf.Port, nil)
	}))
}
