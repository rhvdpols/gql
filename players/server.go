package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rhvdpols/gql/players/graph"
	"github.com/rhvdpols/gql/players/graph/model"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	players := []*model.Player{
		{ID: "1", Name: "Wim", Profile: &model.Profile{Bio: "I'm pretty good at Mario Kart"}},
		{ID: "2", Name: "Ron", Profile: &model.Profile{Bio: "I'm pretty medium at games"}},
		{ID: "3", Name: "Tim", Profile: &model.Profile{Bio: "I'm just not as good at Mario Kart as Wim :-("}},
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Players: players,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
