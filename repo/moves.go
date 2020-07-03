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

func (r *MovesRepo) PerformMove(i int, j int, player *model.Player, game *model.Game) (*model.Move, error) {
	if !game.Started {
		return nil, errors.New("Game hasn't started yet!")
	}

	if game.CurrentTurnID != player.ID {
		return nil, errors.New("It's not your turn!")
	}

	for _, move := range r.moves {
		if move.I == i && move.J == j && move.GameID == game.ID {
			return nil, errors.New("Couldn't perform move.")
		}
	}

	move := &model.Move{
		ID:       uuid.New().String(),
		I:        i,
		J:        j,
		PlayerID: player.ID,
		GameID:   game.ID,
	}

	r.moves[move.ID] = move

	return move, nil
}
