//go:generate go run github.com/99designs/gqlgen
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/seklyza/tictactoe-server/auth"
	"github.com/seklyza/tictactoe-server/graph"
	"github.com/seklyza/tictactoe-server/graph/generated"
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

	resolver := graph.CreateResolver()

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	srv.AddTransport(transport.POST{})
	srv.Use(handler.OperationFunc(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		if debug {
			graphql.GetOperationContext(ctx).DisableIntrospection = false
		} else {
			graphql.GetOperationContext(ctx).DisableIntrospection = true
		}
		return next(ctx)
	}))
	auth.AddWSAuthTransport(srv, resolver.Repos)

	if debug {
		router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	}
	router.Group(func(r chi.Router) {
		r.Use(auth.Middleware(resolver.Repos))
		r.Handle("/graphql", srv)

	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
