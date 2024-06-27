package repositories

import (
	"myapp/internal/entities"
	"time"

	"gorm.io/gorm"
)

type CreatePostRepository struct {
	Conn *gorm.DB
}

// This struct is same as entity model
// However define again for training
type post struct {
	Id        int
	ChannelId int
	UserId    int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewCreatePostRepository(conn *gorm.DB) *CreatePostRepository {
	return &CreatePostRepository{
		Conn: conn,
	}
}

func (r *CreatePostRepository) Post(userID int, channelID int, content string) (*entities.Post, error) {
	var post post
	post.ChannelId = channelID
	post.UserId = userID
	post.Content = content
	result := r.Conn.Omit("DeletedAt").Create(&post)

	if result.Error != nil {
		return nil, result.Error
	}
	return convertCreatePostRepositoryModelToEntity(&post), nil
}

func convertCreatePostRepositoryModelToEntity(v *post) *entities.Post {
	return &entities.Post{
		Id:        v.Id,
		ChannelId: v.ChannelId,
		User:      entities.User{Id: v.UserId},
		Content:   v.Content,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		DeletedAt: v.DeletedAt,
	}
}
