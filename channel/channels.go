package channel

import "github.com/seklyza/tictactoe-server/model"

type Channels struct {
	GameStarts map[string]chan *model.Game
}

func CreateChannels() *Channels {
	return &Channels{
		GameStarts: map[string]chan *model.Game{},
	}
}
