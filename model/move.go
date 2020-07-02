package model

type Move struct {
	ID       string `json:"id"`
	Index    int    `json:"index"`
	PlayerID string `json:"playerId"`
	GameID   string `json:"gameId"`
}
