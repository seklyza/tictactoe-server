package auth

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/seklyza/tictactoe-server/model"
	"github.com/seklyza/tictactoe-server/repo"
)

const currentPlayerKey = "currentPlayer"

func keyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}

func Middleware(repos *repo.Repos) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, keyFunc)

			if err != nil || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			claims := token.Claims.(jwt.MapClaims)

			id := claims["id"].(string)

			player, err := repos.PlayersRepo.GetPlayerByID(id)

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), currentPlayerKey, player) // nolint

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func GetPlayerFromToken(repos *repo.Repos, tokenString string) (*model.Player, error) {
	token, err := jwt.Parse(tokenString, keyFunc)

	if err != nil || !token.Valid {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	player, err := repos.PlayersRepo.GetPlayerByID(id)

	if err != nil {
		return nil, err
	}

	return player, nil
}

func AddWSAuthTransport(srv *handler.Server, repos *repo.Repos) {
	srv.AddTransport(transport.Websocket{InitFunc: func(ctx context.Context, payload transport.InitPayload) (context.Context, error) {
		token, err := request.AuthorizationHeaderExtractor.Filter(payload.Authorization())
		if err != nil {
			return ctx, nil
		}

		player, err := GetPlayerFromToken(repos, token)
		if err != nil {
			return ctx, nil
		}

		ctx = context.WithValue(ctx, currentPlayerKey, player) // nolint

		return ctx, nil
	}})
}

func GetCurrentPlayer(ctx context.Context) (*model.Player, error) {
	value, ok := ctx.Value(currentPlayerKey).(*model.Player)

	if value == nil || !ok {
		return nil, errors.New("Unauthenticated")
	}

	return value, nil
}
