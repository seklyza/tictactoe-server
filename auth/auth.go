package auth

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/seklyza/tictactoe-server/model"
	"github.com/seklyza/tictactoe-server/repo"
)

const currentPlayerKey = "currentPlayer"

func Middleware(repos *repo.Repos) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

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

func GetCurrentPlayer(ctx context.Context) (*model.Player, error) {
	value, ok := ctx.Value(currentPlayerKey).(*model.Player)

	if value == nil || !ok {
		return nil, errors.New("Unauthenticated")
	}

	return value, nil
}
