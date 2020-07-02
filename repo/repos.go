package repo

import "github.com/seklyza/tictactoe-server/model"

type Repos struct {
	GamesRepo   *GamesRepo
	PlayersRepo *PlayersRepo
}

func CreateRepos() *Repos {
	return &Repos{
		GamesRepo: &GamesRepo{
			games: make(map[string]*model.Game),
		},
		PlayersRepo: &PlayersRepo{
			players: make(map[string]*model.Player),
		},
	}
}
