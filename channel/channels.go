package channel

import "github.com/seklyza/tictactoe-server/model"

type Channels struct {
	GameStarts map[string]chan *model.Game
	NewMove    map[string]chan *model.Move
}

func CreateChannels() *Channels {
	return &Channels{
		GameStarts: map[string]chan *model.Game{},
		NewMove:    map[string]chan *model.Move{},
	}
}

func (r *Channels) MakeChannelsForGame(gameId string) {
	r.GameStarts[gameId] = make(chan *model.Game)
	r.NewMove[gameId] = make(chan *model.Move)
}
