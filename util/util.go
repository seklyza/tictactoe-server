package util

import (
	"crypto/rand"
	"io"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/seklyza/tictactoe-server/model"
)

var table = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func GenerateGameCode() string {
	max := 6
	b := make([]byte, max)
	_, _ = io.ReadAtLeast(rand.Reader, b, max)
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func GenerateToken(id string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
}

/*
   0 1 2   - - -   - - -   0 - -   - 1 -   - - 2   0 - -   - - 2
   - - -   3 4 5   - - -   3 - -   - 4 -   - - 5   - 4 -   - 4 -
   - - -   - - -   6 7 8   6 - -   - 7 -   - - 8   - - 8   6 - -
*/
var winTable = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

func CalculateWinner(moves []*model.Move, gameId string) string {
	board := make([]string, 9)
	for _, move := range moves {
		board[move.J+3*move.I] = move.PlayerID
	}

	for _, winSituation := range winTable {
		if board[winSituation[0]] != "" && board[winSituation[0]] == board[winSituation[1]] && board[winSituation[1]] == board[winSituation[2]] {
			return board[winSituation[0]]
		}
	}

	return ""
}
