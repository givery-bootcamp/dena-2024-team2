package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"myapp/internal/entities"
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

func NewChannelsRepository(conn *gorm.DB) *GetChannelsRepository {
	return &GetChannelsRepository{
		Conn: conn,
	}
}

func (r *GetChannelsRepository) Get() ([]entities.Channel, error) {
	obj := []Channel{}
	r.Conn.Find(&obj)
	fmt.Printf("result: %+v\n", obj)
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