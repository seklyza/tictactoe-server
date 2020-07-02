package repo

import (
	"errors"

	"github.com/google/uuid"
	"github.com/seklyza/tictactoe-server/model"
)

type MovesRepo struct {
	moves map[string]*model.Move
}

func (r *MovesRepo) GetMovesByGameID(gameId string) []*model.Move {
	moves := []*model.Move{}
	for _, move := range r.moves {
		if move.GameID == gameId {
			moves = append(moves, move)
		}
	}

	return moves
}

func (r *MovesRepo) PerformMove(index int, player *model.Player, game *model.Game) (*model.Move, error) {
	if !game.Started {
		return nil, errors.New("Game hasn't started yet!")
	}

	for _, move := range r.moves {
		if move.Index == index {
			return nil, errors.New("Couldn't perform move.")
		}
	}

	move := &model.Move{
		ID:       uuid.New().String(),
		Index:    index,
		PlayerID: player.ID,
		GameID:   game.ID,
	}

	r.moves[move.ID] = move

	return move, nil
}
