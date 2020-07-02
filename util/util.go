package util

import (
	"crypto/rand"
	"io"
	"os"

	"github.com/dgrijalva/jwt-go"
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
