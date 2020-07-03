//go:generate go run github.com/99designs/gqlgen
package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/seklyza/tictactoe-server/auth"
	"github.com/seklyza/tictactoe-server/channel"
	"github.com/seklyza/tictactoe-server/graph"
	"github.com/seklyza/tictactoe-server/graph/generated"
	"github.com/seklyza/tictactoe-server/repo"
)

const defaultPort = "8080"

func main() {
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	repos := repo.CreateRepos()
	channels := channel.CreateChannels()

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Repos:    repos,
		Channels: channels,
	}}))

	srv.AddTransport(transport.POST{})
	srv.Use(handler.OperationFunc(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		graphql.GetOperationContext(ctx).DisableIntrospection = false
		return next(ctx)
	}))
	auth.AddWSAuthTransport(srv, repos)

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Group(func(r chi.Router) {
		r.Use(auth.Middleware(repos))
		r.Handle("/graphql", srv)

	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
