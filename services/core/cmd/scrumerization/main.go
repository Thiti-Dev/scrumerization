package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Thiti-Dev/scrumerization-core-service/graph"
	"github.com/Thiti-Dev/scrumerization-core-service/internal/infrastructure/db/postgres/repositories"
	. "github.com/Thiti-Dev/scrumerization-core-service/internal/infrastructure/utils"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := sql.Open("postgres", config.DBSource)
	panicOnError(err)

	userRepository := repositories.NewUserRepository(db)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		SqlConnection:  db,
		UserRepository: userRepository,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
