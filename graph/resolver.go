package graph

import (
	"sync"

	"github.com/seklyza/tictactoe-server/channel"
	"github.com/seklyza/tictactoe-server/repo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	mu       *sync.Mutex
	Repos    *repo.Repos
	Channels *channel.Channels
}

func CreateResolver() *Resolver {
	return &Resolver{
		mu:       &sync.Mutex{},
		Repos:    repo.CreateRepos(),
		Channels: channel.CreateChannels(),
	}
}
