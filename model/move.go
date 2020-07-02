package model

type Move struct {
	ID       string `json:"id"`
	PlayerID string `json:"playerId"`
	Index    int    `json:"index"`
}
