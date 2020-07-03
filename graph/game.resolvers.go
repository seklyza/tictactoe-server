package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/seklyza/tictactoe-server/auth"
	"github.com/seklyza/tictactoe-server/graph/generated"
	"github.com/seklyza/tictactoe-server/model"
	"github.com/seklyza/tictactoe-server/util"
)

func (r *gameResolver) CurrentTurn(ctx context.Context, obj *model.Game) (*model.Player, error) {
	return r.Repos.PlayersRepo.GetPlayerByID(obj.CurrentTurnID)
}

func (r *gameResolver) Players(ctx context.Context, obj *model.Game) ([]*model.Player, error) {
	return r.Repos.PlayersRepo.GetPlayersByGameID(obj.ID), nil
}

func (r *gameResolver) Moves(ctx context.Context, obj *model.Game) ([]*model.Move, error) {
	return r.Repos.MovesRepo.GetMovesByGameID(obj.ID), nil
}

func (r *moveResolver) Player(ctx context.Context, obj *model.Move) (*model.Player, error) {
	return r.Repos.PlayersRepo.GetPlayerByID(obj.PlayerID)
}

func (r *moveResolver) Game(ctx context.Context, obj *model.Move) (*model.Game, error) {
	return r.Repos.GamesRepo.GetGameByID(obj.GameID)
}

func (r *mutationResolver) CreateGame(ctx context.Context) (*model.JoinGameResult, error) {
	game := r.Repos.GamesRepo.CreateGame()

	me := r.Repos.PlayersRepo.CreatePlayer(model.PlayerTypeX, game.ID)
	game.CurrentTurnID = me.ID

	token, err := util.GenerateToken(me.ID)

	if err != nil {
		return nil, err
	}

	r.Channels.MakeChannelsForGame(game.ID)

	return &model.JoinGameResult{
		Game:  game,
		Token: token,
	}, nil
}

func (r *mutationResolver) JoinGame(ctx context.Context, code string) (*model.JoinGameResult, error) {
	game, err := r.Repos.GamesRepo.GetGameByCode(code)

	if err != nil {
		return nil, err
	}

	if game.Started {
		return nil, errors.New("Game has already started.")
	}

	me := r.Repos.PlayersRepo.CreatePlayer(model.PlayerTypeO, game.ID)
	token, err := util.GenerateToken(me.ID)

	if err != nil {
		return nil, err
	}

	game.Started = true
	for _, c := range r.Channels.GameStarts[game.ID] {
		go func(c chan *model.Game) {
			c <- game
		}(c)
	}

	return &model.JoinGameResult{
		Game:  game,
		Token: token,
	}, nil
}

func (r *mutationResolver) PerformMove(ctx context.Context, i int, j int) (*model.Move, error) {
	me, err := auth.GetCurrentPlayer(ctx)

	if err != nil {
		return nil, err
	}

	game, err := r.Repos.GamesRepo.GetGameByID(me.GameID)

	if err != nil {
		return nil, err
	}

	move, err := r.Repos.MovesRepo.PerformMove(i, j, me, game)

	if err != nil {
		return nil, err
	}

	for _, player := range r.Repos.PlayersRepo.GetPlayersByGameID(me.GameID) {
		if player.ID != me.ID {
			game.CurrentTurnID = player.ID
			break
		}
	}

	for _, c := range r.Channels.NewMove[game.ID] {
		go func(c chan *model.Move) {
			c <- move
		}(c)
	}

	return move, nil
}

func (r *playerResolver) Moves(ctx context.Context, obj *model.Player) ([]*model.Move, error) {
	return r.Repos.MovesRepo.GetMovesByGameID(obj.GameID), nil
}

func (r *playerResolver) Game(ctx context.Context, obj *model.Player) (*model.Game, error) {
	return r.Repos.GamesRepo.GetGameByID(obj.GameID)
}

func (r *queryResolver) Me(ctx context.Context) (*model.Player, error) {
	me, err := auth.GetCurrentPlayer(ctx)

	if err != nil {
		return nil, nil
	}

	return me, nil
}

func (r *subscriptionResolver) GameStarts(ctx context.Context) (<-chan *model.Game, error) {
	player, err := auth.GetCurrentPlayer(ctx)

	if err != nil {
		return nil, err
	}

	r.Channels.GameStarts[player.GameID][player.ID] = make(chan *model.Game)

	return r.Channels.GameStarts[player.GameID][player.ID], nil
}

func (r *subscriptionResolver) NewMove(ctx context.Context) (<-chan *model.Move, error) {
	player, err := auth.GetCurrentPlayer(ctx)

	if err != nil {
		return nil, err
	}

	r.Channels.NewMove[player.GameID][player.ID] = make(chan *model.Move, 9)

	return r.Channels.NewMove[player.GameID][player.ID], nil
}

func (r *subscriptionResolver) GameEnds(ctx context.Context) (<-chan *model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

// Move returns generated.MoveResolver implementation.
func (r *Resolver) Move() generated.MoveResolver { return &moveResolver{r} }

// Player returns generated.PlayerResolver implementation.
func (r *Resolver) Player() generated.PlayerResolver { return &playerResolver{r} }

type gameResolver struct{ *Resolver }
type moveResolver struct{ *Resolver }
type playerResolver struct{ *Resolver }
