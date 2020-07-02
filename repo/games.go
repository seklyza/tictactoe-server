package repo

import (
	"errors"

	"github.com/google/uuid"
	"github.com/seklyza/tictactoe-server/model"
	"github.com/seklyza/tictactoe-server/util"
)

type GamesRepo struct {
	games map[string]*model.Game
}

func (r *GamesRepo) CreateGame() *model.Game {
	id := uuid.New().String()
	code := util.GenerateGameCode()
	game := &model.Game{
		ID:   id,
		Code: code,
	}

	r.games[game.ID] = game

	return game
}

func (r *GamesRepo) GetGameByID(id string) (*model.Game, error) {
	game, ok := r.games[id]

	if !ok {
		return nil, errors.New("Game not found.")
	}

	return game, nil
}

func (r *GamesRepo) GetGameByCode(code string) (*model.Game, error) {
	var game *model.Game
	for _, g := range r.games {
		if g.Code == code {
			game = g
			break
		}
	}

	if game == nil {
		return nil, errors.New("Game does not exist.")
	}

	return game, nil
}
