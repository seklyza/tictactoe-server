package repo

import (
	"errors"

	"github.com/google/uuid"
	"github.com/seklyza/tictactoe-server/model"
)

type PlayersRepo struct {
	players map[string]*model.Player
}

func (r *PlayersRepo) CreatePlayer(playerType model.PlayerType, gameId string) *model.Player {
	player := &model.Player{
		ID:         uuid.New().String(),
		GameID:     gameId,
		PlayerType: playerType,
	}

	r.players[player.ID] = player

	return player
}

func (r *PlayersRepo) GetPlayerByID(id string) (*model.Player, error) {
	player, ok := r.players[id]

	if !ok {
		return nil, errors.New("Could not find player.")
	}

	return player, nil
}

func (r *PlayersRepo) GetPlayersByGameID(gameId string) []*model.Player {
	players := []*model.Player{}

	for _, player := range r.players {
		if player.GameID == gameId {
			players = append(players, player)
		}
	}

	return players
}
