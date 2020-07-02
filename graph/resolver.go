package graph

import "github.com/seklyza/tictactoe-server/repo"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repos *repo.Repos
}
