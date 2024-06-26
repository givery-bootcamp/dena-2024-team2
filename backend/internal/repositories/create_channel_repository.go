package repositories

import (
	"fmt"
	"myapp/internal/entities"
	"time"

	"gorm.io/gorm"
)

type CreateChannelRepository struct {
	Conn *gorm.DB
}

type createChannel struct {
	Id        int
	ServerId  int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewCreateChannelRepository(conn *gorm.DB) *CreateChannelRepository {
	return &CreateChannelRepository{
		Conn: conn,
	}
}

func (r *CreateChannelRepository) Post(serverId int, name string) (*entities.Channel, error) {

	channel := createChannel{}
	r.Conn.Table("channels").Select("name").Where("name = ?", name).Find(&channel)

	if channel.Name != "" {
		// name がすでに使われている
		return nil, fmt.Errorf("%v", "その名前はすでに使用されています。")
	}

	// 新しいチャンネルのとき保存する
	newChannel := createChannel{ServerId: serverId, Name: name}
	r.Conn.Table("channels").Omit("DeletedAt").Create(&newChannel)

	return convertCreateChannelRepositoryModelToEntity(&newChannel), nil
}

func convertCreateChannelRepositoryModelToEntity(v *createChannel) *entities.Channel {
	return &entities.Channel{
		Id:        v.Id,
		ServerId:  v.ServerId,
		Name:      v.Name,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		DeletedAt: v.DeletedAt,
	}
}
