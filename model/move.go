package model

type Move struct {
	ID       string `json:"id"`
	I        int    `json:"i"`
	J        int    `json:"j"`
	PlayerID string `json:"playerId"`
	GameID   string `json:"gameId"`
}
