package graph

import "github.com/rhvdpols/gql/games/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Games []*model.Game
}
