package interfaces

import "myapp/internal/entities"

type ServerRepository interface {
	Get(serverId int) (*entities.Server, error)
}
