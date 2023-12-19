package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tabed23/auth-graphql/config"
	"github.com/tabed23/auth-graphql/graph"
	"github.com/tabed23/auth-graphql/middleware"
	"github.com/tabed23/auth-graphql/repository/store"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	database := config.ConnectDB()

	r := mux.NewRouter()
	r.Use(middleware.AuthMiddleware)

	store := store.NewUserStore(database)

	c := graph.Config{Resolvers: &graph.Resolver{UserRepository: store}}
	c.Directives.Auth = middleware.Auth

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
