package main

import (
	"log"
	"net/http"

	"go-crypto-tracker/internal/gql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

// CORS middleware'i ekleme
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := defaultPort

	srv := handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{
		Resolvers: &gql.Resolver{},
	}))

	// CORS middleware'ini GraphQL handler'ınıza ekleyin
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", enableCors(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
