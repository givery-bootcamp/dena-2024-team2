package repositories

import (
	"myapp/internal/entities"
	"time"

	"gorm.io/gorm"
)

type GetChannelsRepository struct {
	Conn *gorm.DB
}

type Channels []*Channel
type Channel struct {
	Id        int
	ServerId  int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewGetChannelsRepository(conn *gorm.DB) *GetChannelsRepository {
	return &GetChannelsRepository{
		Conn: conn,
	}
}

func (r *GetChannelsRepository) Get(serverId int) ([]entities.Channel, error) {

	obj := []Channel{}
	if err := r.Conn.Table("channels").
		Select("id, server_id, name, created_at, updated_at").
		Where("server_id = ?", serverId).
		Where("deleted_at IS NULL").
		Find(&obj).Error; err != nil {
		return nil, err
	}
	return convertChannelsRepositoryModelToEntity(obj), nil
}

func convertChannelsRepositoryModelToEntity(channels []Channel) []entities.Channel {
	entityChannels := make([]entities.Channel, len(channels))
	for i, v := range channels {
		entityChannels[i] = entities.Channel{
			Id:        v.Id,
			ServerId:  v.ServerId,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		}
	}
	return entityChannels
}
