package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rhvdpols/gql/games/graph"
	"github.com/rhvdpols/gql/games/graph/model"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Games: []*model.Game{{ID: "1", Name: "Game 1"},
		{ID: "2", Name: "MARIO KART"},
		{ID: "3", Name: "Game 3"},
		{ID: "4", Name: "Game 4"},
		{ID: "5", Name: "Game 5"},
		{ID: "6", Name: "Game 6"}}}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
