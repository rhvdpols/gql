package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rhvdpols/gql/monolith/graph"
	"github.com/rhvdpols/gql/monolith/graph/model"
)

const defaultPort = "8083"

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

	leaderboards := []*model.Leaderboard{
		{Rankings: []*model.Ranking{
			{Rank: 1, Player: players[0], Points: 1000,},
			{Rank: 2, Player: players[1], Points: 900,},
			{Rank: 3, Player: players[2], Points: 1,},
		}},
		{Rankings: []*model.Ranking{
			{Rank: 1, Player: players[2], Points: 1000,},
			{Rank: 2, Player: players[1], Points: 900,},
			{Rank: 3, Player: players[0], Points: 1,},
		}},
	}

	games := []*model.Game{
		{ID: "1", Name: "Mario Kart", Leaderboard: leaderboards[0]},
		{ID: "2", Name: "Super Mario Bros", Leaderboard: leaderboards[1]},
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Games: games,
		LeaderBoards: leaderboards,
		Players: players,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
