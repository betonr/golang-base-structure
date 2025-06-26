package main

import (
	"log"
	"net/http"

	"github.com/betonr/golang-base-structure/graph"
	"github.com/betonr/golang-base-structure/internal/database/memory"
	"github.com/betonr/golang-base-structure/internal/service"
	"github.com/betonr/golang-base-structure/pkg/logger"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	repo := memory.NewMemoryArticleRepository()
	svc := service.NewArticleService(repo)
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{ArticleService: svc}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.Use(extension.Introspection{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Info("Servidor GraphQL iniciado na porta %d", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
