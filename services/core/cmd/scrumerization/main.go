package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Thiti-Dev/scrumerization-core-service/cmd/scrumerization/wrappers"
	"github.com/Thiti-Dev/scrumerization-core-service/cmd/scrumerization/wrappers/middlewares"
	"github.com/Thiti-Dev/scrumerization-core-service/graph"
	"github.com/Thiti-Dev/scrumerization-core-service/internal/domain/rooms"
	"github.com/Thiti-Dev/scrumerization-core-service/internal/infrastructure/db/postgres/repositories"
	infraUtils "github.com/Thiti-Dev/scrumerization-core-service/internal/infrastructure/utils"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config, err := infraUtils.LoadConfig(".")
	if err != nil {
		log.Println("only use config from system environment as no config file is found:", err)
	}

	log.Println("Database endpoint: ", config.DBSource)

	db, err := sql.Open("postgres", config.DBSource)
	panicOnError(err)

	userRepository := repositories.NewUserRepository(db, &config)
	roomRepository := repositories.NewRoomRepository(db, &config)
	topicRepository := repositories.NewTopicRepository(db, &config)
	roomHub := rooms.NewRoomHub()
	c := graph.Config{Resolvers: &graph.Resolver{
		SqlConnection:   db,
		UserRepository:  userRepository,
		RoomRepository:  roomRepository,
		RoomHub:         roomHub,
		TopicRepository: topicRepository,
		Config:          &config,
	}}

	wrappers.RegisterDirectives(&c, &config)

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}).Handler)

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	srv := handler.New(graph.NewExecutableSchema(c)) // calling new instead of NewDefaultServer because we need more granular control of each Transportation
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{}) // for types inspect

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", middlewares.AuthHandlerForGraphql(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
