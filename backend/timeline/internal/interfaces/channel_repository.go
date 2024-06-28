package interfaces

import "myapp/internal/entities"

type ChannelRepository interface {
	Get(channelId int) (*entities.Channel, error)
}
