package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rhvdpols/gql/leaderboards/graph"
	"github.com/rhvdpols/gql/leaderboards/graph/model"
)

const defaultPort = "8082"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	leaderboards := []*model.Leaderboard{
		{Rankings: []*model.Ranking{
			{Rank: 1, Player: &model.Player{ID: "1"}, Points: 1000},
			{Rank: 2, Player: &model.Player{ID: "2"}, Points: 500},
			{Rank: 3, Player: &model.Player{ID: "3"}, Points: 1},
		}},
		{Rankings: []*model.Ranking{
			{Rank: 1, Player: &model.Player{ID: "3"}, Points: 1000},
			{Rank: 2, Player: &model.Player{ID: "2"}, Points: 500},
			{Rank: 3, Player: &model.Player{ID: "1"}, Points: 1},
		}},
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		LeaderBoards: leaderboards,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
