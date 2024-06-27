package interfaces

import "myapp/internal/entities"

type GetChannelsRepository interface {
	Get(serverId int) ([]entities.Channel, error)
}
