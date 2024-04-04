package graph

import "github.com/rhvdpols/gql/leaderboards/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	LeaderBoards []*model.Leaderboard
}
