package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/seklyza/tictactoe-server/auth"
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Repos: repos}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Group(func(r chi.Router) {
		r.Use(auth.Middleware(repos))
		r.Handle("/graphql", srv)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
