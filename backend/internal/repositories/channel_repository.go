package repositories

import (
	"errors"
	"log"
	"myapp/internal/entities"

	"gorm.io/gorm"
)

type ChannelRepository struct {
	Conn *gorm.DB
}

func NewGetChannelRepository(conn *gorm.DB) *ChannelRepository {
	return &ChannelRepository{
		Conn: conn,
	}
}
func (r *ChannelRepository) Get(channelId int) (*entities.Channel, error) {
	channel := &entities.Channel{}
	if err := r.Conn.First(&channel, channelId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entities.ErrNotFound
		}
		log.Printf("%v", err)
		return nil, err
	}
	return channel, nil
}
