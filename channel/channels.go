package channel

import "github.com/seklyza/tictactoe-server/model"

type Channels struct {
	GameStarts map[string]map[string]chan *model.Game
	NewMove    map[string]map[string]chan *model.Move
	GameEnds   map[string]map[string]chan *model.Player
}

func CreateChannels() *Channels {
	return &Channels{
		GameStarts: map[string]map[string]chan *model.Game{},
		NewMove:    map[string]map[string]chan *model.Move{},
		GameEnds:   map[string]map[string]chan *model.Player{},
	}
}

func (r *Channels) MakeChannelsForGame(gameId string) {
	r.GameStarts[gameId] = make(map[string]chan *model.Game)
	r.NewMove[gameId] = make(map[string]chan *model.Move)
	r.GameEnds[gameId] = make(map[string]chan *model.Player)
}
