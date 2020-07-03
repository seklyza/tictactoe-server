package repo

import (
	"errors"

	"github.com/google/uuid"
	"github.com/seklyza/tictactoe-server/model"
	"github.com/seklyza/tictactoe-server/util"
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

func (r *MovesRepo) PerformMove(i int, j int, player *model.Player, game *model.Game) (*model.Move, string, bool, error) {
	if !game.Started || game.Ended {
		return nil, "", false, errors.New("Game hasn't started yet!")
	}

	if game.CurrentTurnID != player.ID {
		return nil, "", false, errors.New("It's not your turn!")
	}

	moves := r.GetMovesByGameID(game.ID)

	for _, move := range moves {
		if move.I == i && move.J == j {
			return nil, "", false, errors.New("Couldn't perform move.")
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

	if winner := util.CalculateWinner(append(moves, move), game.ID); winner != "" {
		game.Ended = true
		return move, winner, false, nil
	}

	if len(moves) >= 8 {
		return nil, "", true, nil
	}

	return move, "", false, nil
}
