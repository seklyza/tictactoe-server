package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/seklyza/tictactoe-server/auth"
	"github.com/seklyza/tictactoe-server/graph/generated"
	"github.com/seklyza/tictactoe-server/model"
	"github.com/seklyza/tictactoe-server/util"
)

func (r *gameResolver) Players(ctx context.Context, obj *model.Game) ([]*model.Player, error) {
	return r.Repos.PlayersRepo.GetPlayersByGameID(obj.ID), nil
}

func (r *gameResolver) Moves(ctx context.Context, obj *model.Game) ([]*model.Move, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *moveResolver) Player(ctx context.Context, obj *model.Move) (*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateGame(ctx context.Context) (*model.JoinGameResult, error) {
	game := r.Repos.GamesRepo.CreateGame()

	me := r.Repos.PlayersRepo.CreatePlayer(model.PlayerTypeX, game.ID)
	token, err := util.GenerateToken(me.ID)

	if err != nil {
		return nil, err
	}

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

	me := r.Repos.PlayersRepo.CreatePlayer(model.PlayerTypeO, game.ID)
	token, err := util.GenerateToken(me.ID)

	if err != nil {
		return nil, err
	}

	game.Started = true

	return &model.JoinGameResult{
		Game:  game,
		Token: token,
	}, nil
}

func (r *playerResolver) Moves(ctx context.Context, obj *model.Player) ([]*model.Move, error) {
	panic(fmt.Errorf("not implemented"))
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

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

// Move returns generated.MoveResolver implementation.
func (r *Resolver) Move() generated.MoveResolver { return &moveResolver{r} }

// Player returns generated.PlayerResolver implementation.
func (r *Resolver) Player() generated.PlayerResolver { return &playerResolver{r} }

type gameResolver struct{ *Resolver }
type moveResolver struct{ *Resolver }
type playerResolver struct{ *Resolver }
