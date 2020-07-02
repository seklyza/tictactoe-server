package model

type Player struct {
	ID         string     `json:"id"`
	GameID     string     `json:"gameId"`
	PlayerType PlayerType `json:"playerType"`
}
