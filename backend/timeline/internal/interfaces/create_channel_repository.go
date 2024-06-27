package interfaces

import "myapp/internal/entities"

type CreateChannelRepository interface {
	Post(serverId int, name string) (*entities.Channel, error)
}
